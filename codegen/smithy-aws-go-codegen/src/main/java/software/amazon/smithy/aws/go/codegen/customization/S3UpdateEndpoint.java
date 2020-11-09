/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://aws.amazon.com/apache2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 *
 *
 */

package software.amazon.smithy.aws.go.codegen.customization;

import java.util.ArrayList;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;
import software.amazon.smithy.aws.traits.ServiceTrait;
import software.amazon.smithy.codegen.core.CodegenException;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.go.codegen.GoCodegenPlugin;
import software.amazon.smithy.go.codegen.GoDelegator;
import software.amazon.smithy.go.codegen.GoSettings;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.SymbolUtils;
import software.amazon.smithy.go.codegen.integration.ConfigField;
import software.amazon.smithy.go.codegen.integration.GoIntegration;
import software.amazon.smithy.go.codegen.integration.MiddlewareRegistrar;
import software.amazon.smithy.go.codegen.integration.RuntimeClientPlugin;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.OperationShape;
import software.amazon.smithy.model.shapes.ServiceShape;
import software.amazon.smithy.model.shapes.ShapeId;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.utils.ListUtils;
import software.amazon.smithy.utils.SetUtils;

/**
 * S3UpdateEndpoint integration serves to apply customizations for S3 service,
 * and modifies the resolved endpoint based on S3 client config or input shape values.
 */
public class S3UpdateEndpoint implements GoIntegration {
    // Middleware name
    private static final String UPDATE_ENDPOINT_INTERNAL_ADDER = "UpdateEndpoint";
    // Middleware options
    private static final String UPDATE_ENDPOINT_INTERNAL_OPTIONS =
            UPDATE_ENDPOINT_INTERNAL_ADDER + "Options";
    private static final String UPDATE_ENDPOINT_INTERNAL_PARAMETER_ACCESSOR =
            UPDATE_ENDPOINT_INTERNAL_ADDER + "ParameterAccessor";

    // s3 shared option constants
    private static final String USE_DUALSTACK_OPTION = "UseDualstack";
    private static final String USE_ARNREGION_OPTION = "UseARNRegion";

    // list of runtime-client plugins
    private final List<RuntimeClientPlugin> runtimeClientPlugins = new ArrayList<>();

    private static boolean isS3Service(Model model, ServiceShape service) {
        return service.expectTrait(ServiceTrait.class).getSdkId().equalsIgnoreCase("S3");
    }

    private static boolean isS3ControlService(Model model, ServiceShape service) {
        return service.expectTrait(ServiceTrait.class).getSdkId().equalsIgnoreCase("S3 Control");
    }

    private static boolean isS3SharedService(Model model, ServiceShape service) {
        return isS3Service(model, service) || isS3ControlService(model, service);
    }

    private static String copyInputFuncName(String inputName) {
        return String.format("copy%sForUpdateEndpoint", inputName);
    }

    private static String getterFuncName(String operationName, String memberName) {
        return String.format("get%s%s", operationName, memberName);
    }

    private static String setterFuncName(String operationName, String memberName) {
        return String.format("set%s%s", operationName, memberName);
    }

    private static String addMiddlewareFuncName(String operationname, String middlewareName) {
        return String.format("add%s%s", operationname, middlewareName);
    }

    /**
     * Gets the sort order of the customization from -128 to 127, with lowest
     * executed first.
     *
     * @return Returns the sort order, defaults to -40.
     */
    @Override
    public byte getOrder() {
        return 127;
    }

    @Override
    public void writeAdditionalFiles(
            GoSettings settings,
            Model model,
            SymbolProvider symbolProvider,
            GoDelegator goDelegator
    ) {
        ServiceShape service = settings.getService(model);

        // if service is s3control
        if (isS3ControlService(model, service)) {
            s3control obj = new s3control(service);
            obj.writeAdditionalFiles(settings, model, symbolProvider, goDelegator);
        }

        // check if service is s3
        if (isS3Service(model, service)) {
            s3 obj = new s3(service);
            obj.writeAdditionalFiles(settings, model, symbolProvider, goDelegator);
        }
    }

    /**
     * Builds the set of runtime plugs used by the presign url customization.
     *
     * @param settings codegen settings
     * @param model    api model
     */
    @Override
    public void processFinalizedModel(GoSettings settings, Model model) {
        ServiceShape service = settings.getService(model);
        for (ShapeId operationId : service.getAllOperations()) {
            final OperationShape operation = model.expectShape(operationId, OperationShape.class);

            // Create a symbol provider because one is not available in this call.
            SymbolProvider symbolProvider = GoCodegenPlugin.createSymbolProvider(model, settings.getModuleName());
            String helperFuncName = addMiddlewareFuncName(
                    symbolProvider.toSymbol(operation).getName(),
                    UPDATE_ENDPOINT_INTERNAL_ADDER
            );

            runtimeClientPlugins.add(RuntimeClientPlugin.builder()
                    .operationPredicate((m,s,o)-> {
                        if (!isS3SharedService(m, s)) {
                            return false;
                        }
                        return o.equals(operation);
                    })
                    .registerMiddleware(MiddlewareRegistrar.builder()
                            .resolvedFunction(SymbolUtils.createValueSymbolBuilder(helperFuncName)
                                    .build())
                            .useClientOptions()
                            .build())
                    .build());
        }
    }

    @Override
    public List<RuntimeClientPlugin> getClientPlugins() {
        runtimeClientPlugins.addAll(ListUtils.of(
                // Add S3 shared config's dualstack option
                RuntimeClientPlugin.builder()
                        .servicePredicate(S3UpdateEndpoint::isS3SharedService)
                        .configFields(ListUtils.of(
                                ConfigField.builder()
                                        .name(USE_DUALSTACK_OPTION)
                                        .type(SymbolUtils.createValueSymbolBuilder("bool")
                                                .putProperty(SymbolUtils.GO_UNIVERSE_TYPE, true)
                                                .build())
                                        .documentation(
                                                "Allows you to enable Dualstack endpoint support for the service.")
                                        .build(),
                                ConfigField.builder()
                                        .name(USE_ARNREGION_OPTION)
                                        .type(SymbolUtils.createValueSymbolBuilder("bool")
                                                .putProperty(SymbolUtils.GO_UNIVERSE_TYPE, true)
                                                .build())
                                        .documentation("Allows you to enable arn region support for the service.")
                                        .build()
                        ))
                        .build()
        ));
        runtimeClientPlugins.addAll(s3.getClientPlugins());
        return runtimeClientPlugins;
    }


    /*
     * s3 class is the private class handling s3 goIntegration for endpoint mutations
     */
    private static class s3 {

        // options to be generated on Client's options type
        private static final String USE_PATH_STYLE_OPTION = "UsePathStyle";
        private static final String USE_ACCELERATE_OPTION = "UseAccelerate";

        // private function getter constant
        private static final String GET_BUCKET_FROM_INPUT = "getBucketFromInput";
        private static final String SUPPORT_ACCELERATE = "supportAccelerate";

        // service shape representing s3
        private final ServiceShape service;
        // list of operations that do not support accelerate
        private final Set<String> NOT_SUPPORT_ACCELERATE = SetUtils.of(
                "ListBuckets", "CreateBucket", "DeleteBucket"
        );

        private s3(ServiceShape service) {
            this.service = service;
        }

        // getClientPlugins returns a list of client plugins for s3 service
        private static List<RuntimeClientPlugin> getClientPlugins() {
            List<RuntimeClientPlugin> list = ListUtils.of(
                    // Add S3 config to use path style host addressing.
                    RuntimeClientPlugin.builder()
                            .servicePredicate(S3UpdateEndpoint::isS3Service)
                            .configFields(ListUtils.of(
                                    ConfigField.builder()
                                            .name(USE_PATH_STYLE_OPTION)
                                            .type(SymbolUtils.createValueSymbolBuilder("bool")
                                                    .putProperty(SymbolUtils.GO_UNIVERSE_TYPE, true)
                                                    .build())
                                            .documentation(
                                                    "Allows you to enable the client to use path-style addressing, "
                                                            + "i.e., `https://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client "
                                                            + "will use virtual hosted bucket addressing when possible"
                                                            + "(`https://BUCKET.s3.amazonaws.com/KEY`).")
                                            .build(),
                                    ConfigField.builder()
                                            .name(USE_ACCELERATE_OPTION)
                                            .type(SymbolUtils.createValueSymbolBuilder("bool")
                                                    .putProperty(SymbolUtils.GO_UNIVERSE_TYPE, true)
                                                    .build())
                                            .documentation("Allows you to enable S3 Accelerate feature. All operations "
                                                    + "compatible with S3 Accelerate will use the accelerate endpoint for "
                                                    + "requests. Requests not compatible will fall back to normal S3 requests. "
                                                    + "The bucket must be enabled for accelerate to be used with S3 client with "
                                                    + "accelerate enabled. If the bucket is not enabled for accelerate an error "
                                                    + "will be returned. The bucket name must be DNS compatible to work "
                                                    + "with accelerate.")
                                            .build()
                            ))
                            .build()
            );
            return list;
        }

        private void writeAdditionalFiles(
                GoSettings settings,
                Model model,
                SymbolProvider symbolProvider,
                GoDelegator goDelegator
        ) {
            goDelegator.useShapeWriter(service, writer -> {
                writeInputGetter(writer, model, symbolProvider, service);
            });

            goDelegator.useShapeWriter(service, writer -> {
                writeAccelerateValidator(writer, model, symbolProvider, service);
            });

            for (ShapeId operationID : service.getAllOperations()) {
                OperationShape operation = model.expectShape(operationID, OperationShape.class);
                goDelegator.useShapeWriter(operation, writer -> {
                    // get input shape from operation
                    StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);
                    // generate update endpoint middleware helper function
                    writeMiddlewareHelper(writer, model, symbolProvider, operation);
                });
            }
        }

        private void writeMiddlewareHelper(
                GoWriter writer, Model model, SymbolProvider symbolProvider, OperationShape operationShape
        ) {
            // imports
            writer.addUseImports(SmithyGoDependency.SMITHY_MIDDLEWARE);

            // operation name
            String operationName = symbolProvider.toSymbol(operationShape).getName();

            writer.openBlock("func $L(stack *middleware.Stack, options Options) error {", "}",
                    addMiddlewareFuncName(symbolProvider.toSymbol(operationShape).getName(),
                            UPDATE_ENDPOINT_INTERNAL_ADDER), () -> {
                        writer.write("return $T(stack, $T{ \n"
                                        + "Accessor : $T{GetBucketFromInput: $L,\n SupportsAccelerate: $L,\n }, \n"
                                        + "UsePathStyle: options.$L,\n "
                                        + "UseAccelerate: options.$L,\n "
                                        + "EndpointResolver: options.EndpointResolver,\n "
                                        + "EndpointResolverOptions: options.EndpointOptions,\n"
                                        + "UseDualstack: options.$L, \n UseARNRegion: options.$L, \n })",
                                SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_ADDER,
                                        AwsCustomGoDependency.S3_CUSTOMIZATION).build(),
                                SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_OPTIONS,
                                        AwsCustomGoDependency.S3_CUSTOMIZATION).build(),
                                SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_PARAMETER_ACCESSOR,
                                        AwsCustomGoDependency.S3_CUSTOMIZATION).build(),
                                GET_BUCKET_FROM_INPUT,
                                SUPPORT_ACCELERATE,
                                USE_PATH_STYLE_OPTION,
                                USE_ACCELERATE_OPTION,
                                USE_DUALSTACK_OPTION,
                                USE_ARNREGION_OPTION
                        );
                    });
            writer.insertTrailingNewline();
        }

        private void writeInputGetter(
                GoWriter writer, Model model, SymbolProvider symbolProvider, ServiceShape service
        ) {
            writer.writeDocs(
                    "getBucketFromInput returns a boolean indicating if the input has a modeled bucket name, " +
                            " and a pointer to string denoting a provided bucket member value");
            writer.openBlock("func getBucketFromInput(input interface{}) (*string, bool) {", "}", () -> {
                writer.openBlock("switch i:= input.(type) {", "}", () -> {
                    service.getAllOperations().forEach((operationId) -> {
                        OperationShape operation = model.expectShape(operationId, OperationShape.class);
                        StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);

                        List<MemberShape> targetBucketShape = input.getAllMembers().values().stream()
                                .filter(m -> m.getTarget().getName().equals("BucketName"))
                                .collect(Collectors.toList());
                        // if model has multiple top level shapes targeting `BucketName`, we throw a codegen exception
                        if (targetBucketShape.size() > 1) {
                            throw new CodegenException(
                                    "BucketName shape should be targeted by only one input member, found " +
                                            targetBucketShape.size() + " for Input shape: " + input.getId());
                        }

                        if (!targetBucketShape.isEmpty() && !operationId.getName().equalsIgnoreCase(
                                "GetBucketLocation")) {
                            writer.write("case $P: return i.$L, true", symbolProvider.toSymbol(input),
                                    targetBucketShape.get(0).getMemberName());
                        }
                    });
                    writer.write("default: return nil, false");
                });
            });
        }

        private void writeAccelerateValidator(
                GoWriter writer, Model model, SymbolProvider symbolProvider, ServiceShape service
        ) {
            writer.writeDocs(
                    "supportAccelerate returns a boolean indicating if the operation associated with the provided input "
                            + "supports S3 Transfer Acceleration");
            writer.openBlock("func $L(input interface{}) bool {", "}", SUPPORT_ACCELERATE, () -> {
                writer.openBlock("switch input.(type) {", "}", () -> {
                    for (ShapeId operationId : service.getAllOperations()) {
                        // check if operation does not support s3 accelerate
                        if (NOT_SUPPORT_ACCELERATE.contains(operationId.getName())) {
                            OperationShape operation = model.expectShape(operationId, OperationShape.class);
                            StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);
                            writer.write("case $P: return false", symbolProvider.toSymbol(input));
                        }
                    }
                    writer.write("default: return true");
                });
            });
        }
    }

    /**
     * s3control class is the private class handling s3control goIntegration for endpoint mutations
     */
    private static class s3control {

        // S3Control function getter constant
        private static final String GET_OUTPOST_ID_FROM_INPUT = "getOutpostIDFromInput";
        // service associated with this class
        private final ServiceShape service;
        // List of operations that use Accesspoint field as ARN input source.
        private final Set<String> LIST_ACCESSPOINT_ARN_INPUT = SetUtils.of(
                "GetAccessPoint", "DeleteAccessPoint", "PutAccessPointPolicy",
                "GetAccessPointPolicy", "DeleteAccessPointPolicy"
        );
        // List of operations that use OutpostID to resolve endpoint
        private final Set<String> LIST_OUTPOST_ID_INPUT = SetUtils.of(
                "CreateBucket", "ListRegionalBuckets"
        );

        private s3control(ServiceShape service) {
            this.service = service;
        }

        // returns a function identifier string for backfillAccountID function
        private static final String backFillAccountIDFuncName(String operation) {
            return String.format("backFill%s%s", operation, "AccountID");
        }

        // returns a function identifier string for arn member setter function
        private static final String setARNMemberFuncName(String operation) {
            return setterFuncName(operation, "ARNMember");
        }

        // returns a function identifier string for arn member getter function
        private static final String getARNMemberFuncName(String operation) {
            return getterFuncName(operation, "ARNMember");
        }

        void writeAdditionalFiles(
                GoSettings settings,
                Model model,
                SymbolProvider symbolProvider,
                GoDelegator goDelegator
        ) {
            goDelegator.useShapeWriter(service, writer -> {
                // generate outpost id helper function
                writeOutpostIDHelper(writer, model, symbolProvider, service);
            });

            for (ShapeId operationID : service.getAllOperations()) {
                OperationShape operation = model.expectShape(operationID, OperationShape.class);
                goDelegator.useShapeWriter(operation, writer -> {
                    // get input shape from operation
                    StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);
                    // generate input copy function
                    writeInputCopy(writer, symbolProvider, input);
                    // generate arn helper function
                    writeARNHelper(writer, model, symbolProvider, operation);
                    // generate backfill account id helper function
                    writeBackfillAccountIDHelper(writer, model, symbolProvider, operation);
                    // generate update endpoint middleware helper function
                    writeMiddlewareHelper(writer, model, symbolProvider, operation);
                });
            }
        }

        private void writeMiddlewareHelper(
                GoWriter writer, Model model, SymbolProvider symbolProvider, OperationShape operationShape
        ) {
            // imports
            writer.addUseImports(SmithyGoDependency.SMITHY_MIDDLEWARE);

            // input shape
            StructureShape inputShape = model.expectShape(operationShape.getInput().get(), StructureShape.class);
            String operationName = symbolProvider.toSymbol(operationShape).getName();

            writer.openBlock("func $L(stack *middleware.Stack, options Options) error {", "}",
                    addMiddlewareFuncName(symbolProvider.toSymbol(operationShape).getName(),
                            UPDATE_ENDPOINT_INTERNAL_ADDER), () -> {
                        writer.write("return $T(stack, $T{ \n"
                                        + "Accessor : $T{GetARNInput: $L,\n BackfillAccountID: $L,\n"
                                        + "GetOutpostIDInput: $L, \n UpdateARNField: $L,\n CopyInput: $L,\n }, \n"
                                        + "EndpointResolver: options.EndpointResolver,\n "
                                        + "EndpointResolverOptions: options.EndpointOptions,\n"
                                        + "UseDualstack: options.$L, \n UseARNRegion: options.$L, \n })",
                                SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_ADDER,
                                        AwsCustomGoDependency.S3CONTROL_CUSTOMIZATION).build(),
                                SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_OPTIONS,
                                        AwsCustomGoDependency.S3CONTROL_CUSTOMIZATION).build(),
                                SymbolUtils.createValueSymbolBuilder(UPDATE_ENDPOINT_INTERNAL_PARAMETER_ACCESSOR,
                                        AwsCustomGoDependency.S3CONTROL_CUSTOMIZATION).build(),
                                LIST_OUTPOST_ID_INPUT.contains(operationName) ? "nil" : getARNMemberFuncName(
                                        operationName),
                                LIST_OUTPOST_ID_INPUT.contains(operationName) ? "nil" : backFillAccountIDFuncName(
                                        operationName),
                                GET_OUTPOST_ID_FROM_INPUT,
                                LIST_OUTPOST_ID_INPUT.contains(operationName) ? "nil" : setARNMemberFuncName(
                                        operationName),
                                copyInputFuncName(symbolProvider.toSymbol(inputShape).getName()),
                                USE_DUALSTACK_OPTION,
                                USE_ARNREGION_OPTION
                        );
                    });
            writer.insertTrailingNewline();
        }

        /**
         * Writes a accessor function that returns an address to copy of passed in input
         *
         * @param writer
         * @param symbolProvider
         * @param input
         */
        private void writeInputCopy(
                GoWriter writer,
                SymbolProvider symbolProvider,
                StructureShape input
        ) {
            Symbol inputSymbol = symbolProvider.toSymbol(input);
            writer.openBlock("func $L(params interface{}) (interface{}, error) {", "}",
                    copyInputFuncName(inputSymbol.getName()),
                    () -> {
                        writer.addUseImports(SmithyGoDependency.FMT);
                        writer.write("input, ok := params.($P)", inputSymbol);
                        writer.openBlock("if !ok {", "}", () -> {
                            writer.write("return nil, fmt.Errorf(\"expect $P type, got %T\", params)", inputSymbol);
                        });
                        writer.write("cpy := *input");
                        writer.write("return &cpy, nil");
                    });
        }

        /**
         * writes BackfillAccountID Helper function for s3 api operation
         * <p>
         * Generates code:
         * === api_operation.go===
         * func backfillAccountID(input interface{}, v string) error {
         * in := input.(*OpInputType)
         * if in.AccountId!=nil {
         * iv := *in.AccountId
         * if !strings.EqualFold(iv, v) {
         * return fmt.Errorf("error backfilling account id")
         * }
         * return nil
         * }
         * <p>
         * in.AccountId = &v
         * return nil
         * }
         */
        private void writeBackfillAccountIDHelper(
                GoWriter writer, Model model, SymbolProvider symbolProvider, OperationShape operation
        ) {
            StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);
            List<MemberShape> targetAccountIDShape = input.getAllMembers().values().stream()
                    .filter(m -> m.getMemberName().equals("AccountId"))
                    .collect(Collectors.toList());
            // if model has multiple top level shapes targeting `AccountId`, we throw a codegen exception
            if (targetAccountIDShape.size() > 1) {
                throw new CodegenException("AccountId shape should be targeted by only one input member, found " +
                        targetAccountIDShape.size() + " for Input shape: " + input.getId());
            }

            Symbol inputSymbol = symbolProvider.toSymbol(input);
            writer.write("func $L (input interface{}, v string) error { ",
                    backFillAccountIDFuncName(symbolProvider.toSymbol(operation).getName()));
            if (!targetAccountIDShape.isEmpty()) {
                String memberName = targetAccountIDShape.get(0).getMemberName();
                writer.write("in := input.($P)", inputSymbol);
                writer.write("if in.$L != nil {", memberName);

                writer.addUseImports(SmithyGoDependency.STRINGS);
                writer.write("if !strings.EqualFold(*in.$L, v) {", memberName);

                writer.addUseImports(SmithyGoDependency.FMT);
                writer.write("return fmt.Errorf(\"error backfilling account id\") }");
                writer.write("return nil }");
                writer.write("in.$L = &v", memberName);
            }
            writer.write("return nil }");
        }

        /**
         * writes getARNMemberValue and updateARNMemberValue update function for all api input operations
         */
        private void writeARNHelper(
                GoWriter writer, Model model, SymbolProvider symbolProvider, OperationShape operation
        ) {
            // list of outpost id input require special behavior
            if (LIST_OUTPOST_ID_INPUT.contains(operation.getId().getName())) {
                return;
            }

            // arn target shape
            String arnType = LIST_ACCESSPOINT_ARN_INPUT.contains(
                    operation.getId().getName()
            ) ? "AccessPointName" : "BucketName";

            StructureShape input = model.expectShape(operation.getInput().get(), StructureShape.class);
            List<MemberShape> listOfARNMembers = input.getAllMembers().values().stream()
                    .filter(m -> m.getTarget().getName().equals(arnType))
                    .collect(Collectors.toList());
            // if model has multiple top level shapes targeting arnable field, we throw a codegen exception
            if (listOfARNMembers.size() > 1) {
                throw new CodegenException(arnType + " shape should be targeted by only one input member, found " +
                        listOfARNMembers.size() + " for Input shape: " + input.getId());
            }

            Symbol inputSymbol = symbolProvider.toSymbol(input);

            // generate arn member accessor getter function
            writer.write("func $L (input interface{}) (*string, bool) { ",
                    getARNMemberFuncName(symbolProvider.toSymbol(operation).getName()));
            if (!listOfARNMembers.isEmpty()) {
                String memberName = listOfARNMembers.get(0).getMemberName();
                writer.write("in := input.($P)", inputSymbol);
                writer.write("if in.$L == nil {return nil, false }", memberName);
                writer.write("return in.$L, true }", memberName);
            } else {
                writer.write("return nil, false }");
            }

            writer.insertTrailingNewline();

            // generate arn member accessor setter function
            writer.write("func $L (input interface{}, v string) error {",
                    setARNMemberFuncName(symbolProvider.toSymbol(operation).getName()));
            if (!listOfARNMembers.isEmpty()) {
                String memberName = listOfARNMembers.get(0).getMemberName();
                writer.write("in := input.($P)", inputSymbol);
                writer.write("in.$L = &v", memberName);
            }
            writer.write("return nil }");
        }

        /**
         * writes OutpostID Helper function for operations CreateBucket and ListRegionalBuckets
         * <p>
         * Generates code:
         * func getOutpostIDFromInput (in interface{}) (*string, bool) {
         * switch in.(type) {
         * case CreateBucket: return in.OutpostId, true
         * case listRegionalBuckets : return in.OutpostID, true
         * default: nil, false
         * }
         * }
         */
        private void writeOutpostIDHelper(
                GoWriter writer, Model model, SymbolProvider symbolProvider, ServiceShape service
        ) {
            writer.writeDocs(
                    "getOutpostIDFromInput returns a boolean indicating if the input has a modeled outpost-id, " +
                            " and a pointer to string denoting a provided outpost-id member value");
            writer.openBlock("func $L (input interface{}) (*string, bool) {", "}",
                    GET_OUTPOST_ID_FROM_INPUT, () -> {
                        writer.openBlock("switch i:= input.(type) {", "}", () -> {
                            for (ShapeId operationId : service.getAllOperations()) {
                                // customization only applied to operations CreateBucket, ListRegionalBuckets
                                if (!LIST_OUTPOST_ID_INPUT.contains(operationId.getName())) {
                                    continue;
                                }
                                OperationShape operation = model.expectShape(operationId, OperationShape.class);
                                StructureShape input = model.expectShape(operation.getInput().get(),
                                        StructureShape.class);
                                List<MemberShape> outpostIDMemberShapes = input.getAllMembers().values().stream()
                                        .filter(m -> m.getMemberName().equalsIgnoreCase("OutpostId"))
                                        .collect(Collectors.toList());
                                // if model has multiple top level shapes targeting `OutpostId`, we throw a codegen exception
                                if (outpostIDMemberShapes.size() > 1) {
                                    throw new CodegenException(
                                            "OutpostID shape should be targeted by only one input member, found " +
                                                    outpostIDMemberShapes.size() + " for Input shape: " + input.getId());
                                }

                                if (!outpostIDMemberShapes.isEmpty()) {
                                    writer.write("case $P: return i.$L, true", symbolProvider.toSymbol(input),
                                            outpostIDMemberShapes.get(0).getMemberName());
                                }
                            }
                            writer.write("default: return nil, false");
                        });
                    });
        }
    }
}
