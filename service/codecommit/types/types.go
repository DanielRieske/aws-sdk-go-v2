// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Returns information about a specific approval on a pull request.
type Approval struct {

	// The state of the approval, APPROVE or REVOKE. REVOKE states are not stored.
	ApprovalState ApprovalState

	// The Amazon Resource Name (ARN) of the user.
	UserArn *string

	noSmithyDocumentSerde
}

// Returns information about an approval rule.
type ApprovalRule struct {

	// The content of the approval rule.
	ApprovalRuleContent *string

	// The system-generated ID of the approval rule.
	ApprovalRuleId *string

	// The name of the approval rule.
	ApprovalRuleName *string

	// The date the approval rule was created, in timestamp format.
	CreationDate *time.Time

	// The date the approval rule was most recently changed, in timestamp format.
	LastModifiedDate *time.Time

	// The Amazon Resource Name (ARN) of the user who made the most recent changes to
	// the approval rule.
	LastModifiedUser *string

	// The approval rule template used to create the rule.
	OriginApprovalRuleTemplate *OriginApprovalRuleTemplate

	// The SHA-256 hash signature for the content of the approval rule.
	RuleContentSha256 *string

	noSmithyDocumentSerde
}

// Returns information about an event for an approval rule.
type ApprovalRuleEventMetadata struct {

	// The content of the approval rule.
	ApprovalRuleContent *string

	// The system-generated ID of the approval rule.
	ApprovalRuleId *string

	// The name of the approval rule.
	ApprovalRuleName *string

	noSmithyDocumentSerde
}

// Returns information about an override event for approval rules for a pull
// request.
type ApprovalRuleOverriddenEventMetadata struct {

	// The status of the override event.
	OverrideStatus OverrideStatus

	// The revision ID of the pull request when the override event occurred.
	RevisionId *string

	noSmithyDocumentSerde
}

// Returns information about an approval rule template.
type ApprovalRuleTemplate struct {

	// The content of the approval rule template.
	ApprovalRuleTemplateContent *string

	// The description of the approval rule template.
	ApprovalRuleTemplateDescription *string

	// The system-generated ID of the approval rule template.
	ApprovalRuleTemplateId *string

	// The name of the approval rule template.
	ApprovalRuleTemplateName *string

	// The date the approval rule template was created, in timestamp format.
	CreationDate *time.Time

	// The date the approval rule template was most recently changed, in timestamp
	// format.
	LastModifiedDate *time.Time

	// The Amazon Resource Name (ARN) of the user who made the most recent changes to
	// the approval rule template.
	LastModifiedUser *string

	// The SHA-256 hash signature for the content of the approval rule template.
	RuleContentSha256 *string

	noSmithyDocumentSerde
}

// Returns information about a change in the approval state for a pull request.
type ApprovalStateChangedEventMetadata struct {

	// The approval status for the pull request.
	ApprovalStatus ApprovalState

	// The revision ID of the pull request when the approval state changed.
	RevisionId *string

	noSmithyDocumentSerde
}

// Returns information about errors in a
// BatchAssociateApprovalRuleTemplateWithRepositories operation.
type BatchAssociateApprovalRuleTemplateWithRepositoriesError struct {

	// An error code that specifies whether the repository name was not valid or not
	// found.
	ErrorCode *string

	// An error message that provides details about why the repository name was not
	// found or not valid.
	ErrorMessage *string

	// The name of the repository where the association was not made.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Returns information about errors in a BatchDescribeMergeConflicts operation.
type BatchDescribeMergeConflictsError struct {

	// The name of the exception.
	//
	// This member is required.
	ExceptionName *string

	// The path to the file.
	//
	// This member is required.
	FilePath *string

	// The message provided by the exception.
	//
	// This member is required.
	Message *string

	noSmithyDocumentSerde
}

// Returns information about errors in a
// BatchDisassociateApprovalRuleTemplateFromRepositories operation.
type BatchDisassociateApprovalRuleTemplateFromRepositoriesError struct {

	// An error code that specifies whether the repository name was not valid or not
	// found.
	ErrorCode *string

	// An error message that provides details about why the repository name was either
	// not found or not valid.
	ErrorMessage *string

	// The name of the repository where the association with the template was not able
	// to be removed.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Returns information about errors in a BatchGetCommits operation.
type BatchGetCommitsError struct {

	// A commit ID that either could not be found or was not in a valid format.
	CommitId *string

	// An error code that specifies whether the commit ID was not valid or not found.
	ErrorCode *string

	// An error message that provides detail about why the commit ID either was not
	// found or was not valid.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// Returns information about errors in a BatchGetRepositories operation.
type BatchGetRepositoriesError struct {

	// An error code that specifies the type of failure.
	ErrorCode BatchGetRepositoriesErrorCodeEnum

	// An error message that provides detail about why the repository either was not
	// found or was not in a valid state.
	ErrorMessage *string

	// The ID of a repository that either could not be found or was not in a valid
	// state.
	RepositoryId *string

	// The name of a repository that either could not be found or was not in a valid
	// state.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Returns information about a specific Git blob object.
type BlobMetadata struct {

	// The full ID of the blob.
	BlobId *string

	// The file mode permissions of the blob. File mode permission codes include:
	//
	//   - 100644 indicates read/write
	//
	//   - 100755 indicates read/write/execute
	//
	//   - 160000 indicates a submodule
	//
	//   - 120000 indicates a symlink
	Mode *string

	// The path to the blob and associated file name, if any.
	Path *string

	noSmithyDocumentSerde
}

// Returns information about a branch.
type BranchInfo struct {

	// The name of the branch.
	BranchName *string

	// The ID of the last commit made to the branch.
	CommitId *string

	noSmithyDocumentSerde
}

// Returns information about a specific comment.
type Comment struct {

	// The Amazon Resource Name (ARN) of the person who posted the comment.
	AuthorArn *string

	// The emoji reactions to a comment, if any, submitted by the user whose
	// credentials are associated with the call to the API.
	CallerReactions []string

	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request is
	// received with the same parameters and a token is included, the request returns
	// information about the initial request that used that token.
	ClientRequestToken *string

	// The system-generated comment ID.
	CommentId *string

	// The content of the comment.
	Content *string

	// The date and time the comment was created, in timestamp format.
	CreationDate *time.Time

	// A Boolean value indicating whether the comment has been deleted.
	Deleted bool

	// The ID of the comment for which this comment is a reply, if any.
	InReplyTo *string

	// The date and time the comment was most recently modified, in timestamp format.
	LastModifiedDate *time.Time

	// A string to integer map that represents the number of individual users who have
	// responded to a comment with the specified reactions.
	ReactionCounts map[string]int32

	noSmithyDocumentSerde
}

// Returns information about comments on the comparison between two commits.
type CommentsForComparedCommit struct {

	// The full blob ID of the commit used to establish the after of the comparison.
	AfterBlobId *string

	// The full commit ID of the commit used to establish the after of the comparison.
	AfterCommitId *string

	// The full blob ID of the commit used to establish the before of the comparison.
	BeforeBlobId *string

	// The full commit ID of the commit used to establish the before of the comparison.
	BeforeCommitId *string

	// An array of comment objects. Each comment object contains information about a
	// comment on the comparison between commits.
	Comments []Comment

	// Location information about the comment on the comparison, including the file
	// name, line number, and whether the version of the file where the comment was
	// made is BEFORE or AFTER.
	Location *Location

	// The name of the repository that contains the compared commits.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Returns information about comments on a pull request.
type CommentsForPullRequest struct {

	// The full blob ID of the file on which you want to comment on the source commit.
	AfterBlobId *string

	// The full commit ID of the commit that was the tip of the source branch at the
	// time the comment was made.
	AfterCommitId *string

	// The full blob ID of the file on which you want to comment on the destination
	// commit.
	BeforeBlobId *string

	// The full commit ID of the commit that was the tip of the destination branch
	// when the pull request was created. This commit is superceded by the after commit
	// in the source branch when and if you merge the source branch into the
	// destination branch.
	BeforeCommitId *string

	// An array of comment objects. Each comment object contains information about a
	// comment on the pull request.
	Comments []Comment

	// Location information about the comment on the pull request, including the file
	// name, line number, and whether the version of the file where the comment was
	// made is BEFORE (destination branch) or AFTER (source branch).
	Location *Location

	// The system-generated ID of the pull request.
	PullRequestId *string

	// The name of the repository that contains the pull request.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Returns information about a specific commit.
type Commit struct {

	// Any other data associated with the specified commit.
	AdditionalData *string

	// Information about the author of the specified commit. Information includes the
	// date in timestamp format with GMT offset, the name of the author, and the email
	// address for the author, as configured in Git.
	Author *UserInfo

	// The full SHA ID of the specified commit.
	CommitId *string

	// Information about the person who committed the specified commit, also known as
	// the committer. Information includes the date in timestamp format with GMT
	// offset, the name of the committer, and the email address for the committer, as
	// configured in Git.
	//
	// For more information about the difference between an author and a committer in
	// Git, see [Viewing the Commit History]in Pro Git by Scott Chacon and Ben Straub.
	//
	// [Viewing the Commit History]: http://git-scm.com/book/ch2-3.html
	Committer *UserInfo

	// The commit message associated with the specified commit.
	Message *string

	// A list of parent commits for the specified commit. Each parent commit ID is the
	// full commit ID.
	Parents []string

	// Tree information for the specified commit.
	TreeId *string

	noSmithyDocumentSerde
}

// Information about conflicts in a merge operation.
type Conflict struct {

	// Metadata about a conflict in a merge operation.
	ConflictMetadata *ConflictMetadata

	// A list of hunks that contain the differences between files or lines causing the
	// conflict.
	MergeHunks []MergeHunk

	noSmithyDocumentSerde
}

// Information about the metadata for a conflict in a merge operation.
type ConflictMetadata struct {

	// A boolean value indicating whether there are conflicts in the content of a file.
	ContentConflict bool

	// A boolean value indicating whether there are conflicts in the file mode of a
	// file.
	FileModeConflict bool

	// The file modes of the file in the source, destination, and base of the merge.
	FileModes *FileModes

	// The path of the file that contains conflicts.
	FilePath *string

	// The file sizes of the file in the source, destination, and base of the merge.
	FileSizes *FileSizes

	// A boolean value (true or false) indicating whether the file is binary or
	// textual in the source, destination, and base of the merge.
	IsBinaryFile *IsBinaryFile

	// Whether an add, modify, or delete operation caused the conflict between the
	// source and destination of the merge.
	MergeOperations *MergeOperations

	// The number of conflicts, including both hunk conflicts and metadata conflicts.
	NumberOfConflicts int32

	// A boolean value (true or false) indicating whether there are conflicts between
	// the branches in the object type of a file, folder, or submodule.
	ObjectTypeConflict bool

	// Information about any object type conflicts in a merge operation.
	ObjectTypes *ObjectTypes

	noSmithyDocumentSerde
}

// If AUTOMERGE is the conflict resolution strategy, a list of inputs to use when
// resolving conflicts during a merge.
type ConflictResolution struct {

	// Files to be deleted as part of the merge conflict resolution.
	DeleteFiles []DeleteFileEntry

	// Files to have content replaced as part of the merge conflict resolution.
	ReplaceContents []ReplaceContentEntry

	// File modes that are set as part of the merge conflict resolution.
	SetFileModes []SetFileModeEntry

	noSmithyDocumentSerde
}

// A file that is deleted as part of a commit.
type DeleteFileEntry struct {

	// The full path of the file to be deleted, including the name of the file.
	//
	// This member is required.
	FilePath *string

	noSmithyDocumentSerde
}

// Returns information about a set of differences for a commit specifier.
type Difference struct {

	// Information about an afterBlob data type object, including the ID, the file
	// mode permission code, and the path.
	AfterBlob *BlobMetadata

	// Information about a beforeBlob data type object, including the ID, the file
	// mode permission code, and the path.
	BeforeBlob *BlobMetadata

	// Whether the change type of the difference is an addition (A), deletion (D), or
	// modification (M).
	ChangeType ChangeTypeEnum

	noSmithyDocumentSerde
}

// Returns information about the approval rules applied to a pull request and
// whether conditions have been met.
type Evaluation struct {

	// The names of the approval rules that have not had their conditions met.
	ApprovalRulesNotSatisfied []string

	// The names of the approval rules that have had their conditions met.
	ApprovalRulesSatisfied []string

	// Whether the state of the pull request is approved.
	Approved bool

	// Whether the approval rule requirements for the pull request have been
	// overridden and no longer need to be met.
	Overridden bool

	noSmithyDocumentSerde
}

// Returns information about a file in a repository.
type File struct {

	// The fully qualified path to the file in the repository.
	AbsolutePath *string

	// The blob ID that contains the file information.
	BlobId *string

	// The extrapolated file mode permissions for the file. Valid values include
	// EXECUTABLE and NORMAL.
	FileMode FileModeTypeEnum

	// The relative path of the file from the folder where the query originated.
	RelativePath *string

	noSmithyDocumentSerde
}

// A file to be added, updated, or deleted as part of a commit.
type FileMetadata struct {

	// The full path to the file to be added or updated, including the name of the
	// file.
	AbsolutePath *string

	// The blob ID that contains the file information.
	BlobId *string

	// The extrapolated file mode permissions for the file. Valid values include
	// EXECUTABLE and NORMAL.
	FileMode FileModeTypeEnum

	noSmithyDocumentSerde
}

// Information about file modes in a merge or pull request.
type FileModes struct {

	// The file mode of a file in the base of a merge or pull request.
	Base FileModeTypeEnum

	// The file mode of a file in the destination of a merge or pull request.
	Destination FileModeTypeEnum

	// The file mode of a file in the source of a merge or pull request.
	Source FileModeTypeEnum

	noSmithyDocumentSerde
}

// Information about the size of files in a merge or pull request.
type FileSizes struct {

	// The size of a file in the base of a merge or pull request.
	Base int64

	// The size of a file in the destination of a merge or pull request.
	Destination int64

	// The size of a file in the source of a merge or pull request.
	Source int64

	noSmithyDocumentSerde
}

// Information about a version of a file.
type FileVersion struct {

	// The blob ID of the object that represents the content of the file in this
	// version.
	BlobId *string

	// Returns information about a specific commit.
	Commit *Commit

	// The name and path of the file at which this blob is indexed which contains the
	// data for this version of the file. This value will vary between file versions if
	// a file is renamed or if its path changes.
	Path *string

	// An array of commit IDs that contain more recent versions of this file. If there
	// are no additional versions of the file, this array will be empty.
	RevisionChildren []string

	noSmithyDocumentSerde
}

// Returns information about a folder in a repository.
type Folder struct {

	// The fully qualified path of the folder in the repository.
	AbsolutePath *string

	// The relative path of the specified folder from the folder where the query
	// originated.
	RelativePath *string

	// The full SHA-1 pointer of the tree information for the commit that contains the
	// folder.
	TreeId *string

	noSmithyDocumentSerde
}

// Information about whether a file is binary or textual in a merge or pull
// request operation.
type IsBinaryFile struct {

	// The binary or non-binary status of a file in the base of a merge or pull
	// request.
	Base *bool

	// The binary or non-binary status of a file in the destination of a merge or pull
	// request.
	Destination *bool

	// The binary or non-binary status of file in the source of a merge or pull
	// request.
	Source *bool

	noSmithyDocumentSerde
}

// Returns information about the location of a change or comment in the comparison
// between two commits or a pull request.
type Location struct {

	// The name of the file being compared, including its extension and subdirectory,
	// if any.
	FilePath *string

	// The position of a change in a compared file, in line number format.
	FilePosition *int64

	// In a comparison of commits or a pull request, whether the change is in the
	// before or after of that comparison.
	RelativeFileVersion RelativeFileVersionEnum

	noSmithyDocumentSerde
}

// Information about merge hunks in a merge or pull request operation.
type MergeHunk struct {

	// Information about the merge hunk in the base of a merge or pull request.
	Base *MergeHunkDetail

	// Information about the merge hunk in the destination of a merge or pull request.
	Destination *MergeHunkDetail

	// A Boolean value indicating whether a combination of hunks contains a conflict.
	// Conflicts occur when the same file or the same lines in a file were modified in
	// both the source and destination of a merge or pull request. Valid values include
	// true, false, and null. True when the hunk represents a conflict and one or more
	// files contains a line conflict. File mode conflicts in a merge do not set this
	// to true.
	IsConflict bool

	// Information about the merge hunk in the source of a merge or pull request.
	Source *MergeHunkDetail

	noSmithyDocumentSerde
}

// Information about the details of a merge hunk that contains a conflict in a
// merge or pull request operation.
type MergeHunkDetail struct {

	// The end position of the hunk in the merge result.
	EndLine *int32

	// The base-64 encoded content of the hunk merged region that might contain a
	// conflict.
	HunkContent *string

	// The start position of the hunk in the merge result.
	StartLine *int32

	noSmithyDocumentSerde
}

// Returns information about a merge or potential merge between a source reference
// and a destination reference in a pull request.
type MergeMetadata struct {

	// A Boolean value indicating whether the merge has been made.
	IsMerged bool

	// The commit ID for the merge commit, if any.
	MergeCommitId *string

	// The merge strategy used in the merge.
	MergeOption MergeOptionTypeEnum

	// The Amazon Resource Name (ARN) of the user who merged the branches.
	MergedBy *string

	noSmithyDocumentSerde
}

// Information about the file operation conflicts in a merge operation.
type MergeOperations struct {

	// The operation on a file in the destination of a merge or pull request.
	Destination ChangeTypeEnum

	// The operation (add, modify, or delete) on a file in the source of a merge or
	// pull request.
	Source ChangeTypeEnum

	noSmithyDocumentSerde
}

// Information about the type of an object in a merge operation.
type ObjectTypes struct {

	// The type of the object in the base commit of the merge.
	Base ObjectTypeEnum

	// The type of the object in the destination branch.
	Destination ObjectTypeEnum

	// The type of the object in the source branch.
	Source ObjectTypeEnum

	noSmithyDocumentSerde
}

// Returns information about the template that created the approval rule for a
// pull request.
type OriginApprovalRuleTemplate struct {

	// The ID of the template that created the approval rule.
	ApprovalRuleTemplateId *string

	// The name of the template that created the approval rule.
	ApprovalRuleTemplateName *string

	noSmithyDocumentSerde
}

// Returns information about a pull request.
type PullRequest struct {

	// The approval rules applied to the pull request.
	ApprovalRules []ApprovalRule

	// The Amazon Resource Name (ARN) of the user who created the pull request.
	AuthorArn *string

	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request is
	// received with the same parameters and a token is included, the request returns
	// information about the initial request that used that token.
	ClientRequestToken *string

	// The date and time the pull request was originally created, in timestamp format.
	CreationDate *time.Time

	// The user-defined description of the pull request. This description can be used
	// to clarify what should be reviewed and other details of the request.
	Description *string

	// The day and time of the last user or system activity on the pull request, in
	// timestamp format.
	LastActivityDate *time.Time

	// The system-generated ID of the pull request.
	PullRequestId *string

	// The status of the pull request. Pull request status can only change from OPEN
	// to CLOSED .
	PullRequestStatus PullRequestStatusEnum

	// The targets of the pull request, including the source branch and destination
	// branch for the pull request.
	PullRequestTargets []PullRequestTarget

	// The system-generated revision ID for the pull request.
	RevisionId *string

	// The user-defined title of the pull request. This title is displayed in the list
	// of pull requests to other repository users.
	Title *string

	noSmithyDocumentSerde
}

// Metadata about the pull request that is used when comparing the pull request
// source with its destination.
type PullRequestCreatedEventMetadata struct {

	// The commit ID of the tip of the branch specified as the destination branch when
	// the pull request was created.
	DestinationCommitId *string

	// The commit ID of the most recent commit that the source branch and the
	// destination branch have in common.
	MergeBase *string

	// The name of the repository where the pull request was created.
	RepositoryName *string

	// The commit ID on the source branch used when the pull request was created.
	SourceCommitId *string

	noSmithyDocumentSerde
}

// Returns information about a pull request event.
type PullRequestEvent struct {

	// The Amazon Resource Name (ARN) of the user whose actions resulted in the event.
	// Examples include updating the pull request with more commits or changing the
	// status of a pull request.
	ActorArn *string

	// Information about a pull request event.
	ApprovalRuleEventMetadata *ApprovalRuleEventMetadata

	// Information about an approval rule override event for a pull request.
	ApprovalRuleOverriddenEventMetadata *ApprovalRuleOverriddenEventMetadata

	// Information about an approval state change for a pull request.
	ApprovalStateChangedEventMetadata *ApprovalStateChangedEventMetadata

	// The day and time of the pull request event, in timestamp format.
	EventDate *time.Time

	// Information about the source and destination branches for the pull request.
	PullRequestCreatedEventMetadata *PullRequestCreatedEventMetadata

	// The type of the pull request event (for example, a status change event
	// (PULL_REQUEST_STATUS_CHANGED) or update event
	// (PULL_REQUEST_SOURCE_REFERENCE_UPDATED)).
	PullRequestEventType PullRequestEventType

	// The system-generated ID of the pull request.
	PullRequestId *string

	// Information about the change in mergability state for the pull request event.
	PullRequestMergedStateChangedEventMetadata *PullRequestMergedStateChangedEventMetadata

	// Information about the updated source branch for the pull request event.
	PullRequestSourceReferenceUpdatedEventMetadata *PullRequestSourceReferenceUpdatedEventMetadata

	// Information about the change in status for the pull request event.
	PullRequestStatusChangedEventMetadata *PullRequestStatusChangedEventMetadata

	noSmithyDocumentSerde
}

// Returns information about the change in the merge state for a pull request
// event.
type PullRequestMergedStateChangedEventMetadata struct {

	// The name of the branch that the pull request is merged into.
	DestinationReference *string

	// Information about the merge state change event.
	MergeMetadata *MergeMetadata

	// The name of the repository where the pull request was created.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Information about an update to the source branch of a pull request.
type PullRequestSourceReferenceUpdatedEventMetadata struct {

	// The full commit ID of the commit in the source branch that was the tip of the
	// branch at the time the pull request was updated.
	AfterCommitId *string

	// The full commit ID of the commit in the destination branch that was the tip of
	// the branch at the time the pull request was updated.
	BeforeCommitId *string

	// The commit ID of the most recent commit that the source branch and the
	// destination branch have in common.
	MergeBase *string

	// The name of the repository where the pull request was updated.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Information about a change to the status of a pull request.
type PullRequestStatusChangedEventMetadata struct {

	// The changed status of the pull request.
	PullRequestStatus PullRequestStatusEnum

	noSmithyDocumentSerde
}

// Returns information about a pull request target.
type PullRequestTarget struct {

	// The full commit ID that is the tip of the destination branch. This is the
	// commit where the pull request was or will be merged.
	DestinationCommit *string

	// The branch of the repository where the pull request changes are merged. Also
	// known as the destination branch.
	DestinationReference *string

	// The commit ID of the most recent commit that the source branch and the
	// destination branch have in common.
	MergeBase *string

	// Returns metadata about the state of the merge, including whether the merge has
	// been made.
	MergeMetadata *MergeMetadata

	// The name of the repository that contains the pull request source and
	// destination branches.
	RepositoryName *string

	// The full commit ID of the tip of the source branch used to create the pull
	// request. If the pull request branch is updated by a push while the pull request
	// is open, the commit ID changes to reflect the new tip of the branch.
	SourceCommit *string

	// The branch of the repository that contains the changes for the pull request.
	// Also known as the source branch.
	SourceReference *string

	noSmithyDocumentSerde
}

// Information about a file added or updated as part of a commit.
type PutFileEntry struct {

	// The full path to the file in the repository, including the name of the file.
	//
	// This member is required.
	FilePath *string

	// The content of the file, if a source file is not specified.
	FileContent []byte

	// The extrapolated file mode permissions for the file. Valid values include
	// EXECUTABLE and NORMAL.
	FileMode FileModeTypeEnum

	// The name and full path of the file that contains the changes you want to make
	// as part of the commit, if you are not providing the file content directly.
	SourceFile *SourceFileSpecifier

	noSmithyDocumentSerde
}

// Information about the reaction values provided by users on a comment.
type ReactionForComment struct {

	// The reaction for a specified comment.
	Reaction *ReactionValueFormats

	// The Amazon Resource Names (ARNs) of users who have provided reactions to the
	// comment.
	ReactionUsers []string

	// A numerical count of users who reacted with the specified emoji whose
	// identities have been subsequently deleted from IAM. While these IAM users or
	// roles no longer exist, the reactions might still appear in total reaction
	// counts.
	ReactionsFromDeletedUsersCount *int32

	noSmithyDocumentSerde
}

// Information about the values for reactions to a comment. CodeCommit supports a
// limited set of reactions.
type ReactionValueFormats struct {

	// The Emoji Version 1.0 graphic of the reaction. These graphics are interpreted
	// slightly differently on different operating systems.
	Emoji *string

	// The emoji short code for the reaction. Short codes are interpreted slightly
	// differently on different operating systems.
	ShortCode *string

	// The Unicode codepoint for the reaction.
	Unicode *string

	noSmithyDocumentSerde
}

// Information about a replacement content entry in the conflict of a merge or
// pull request operation.
type ReplaceContentEntry struct {

	// The path of the conflicting file.
	//
	// This member is required.
	FilePath *string

	// The replacement type to use when determining how to resolve the conflict.
	//
	// This member is required.
	ReplacementType ReplacementTypeEnum

	// The base-64 encoded content to use when the replacement type is USE_NEW_CONTENT.
	Content []byte

	// The file mode to apply during conflict resoltion.
	FileMode FileModeTypeEnum

	noSmithyDocumentSerde
}

// Information about a repository.
type RepositoryMetadata struct {

	// The ID of the Amazon Web Services account associated with the repository.
	AccountId *string

	// The Amazon Resource Name (ARN) of the repository.
	Arn *string

	// The URL to use for cloning the repository over HTTPS.
	CloneUrlHttp *string

	// The URL to use for cloning the repository over SSH.
	CloneUrlSsh *string

	// The date and time the repository was created, in timestamp format.
	CreationDate *time.Time

	// The repository's default branch name.
	DefaultBranch *string

	// The ID of the Key Management Service encryption key used to encrypt and decrypt
	// the repository.
	KmsKeyId *string

	// The date and time the repository was last modified, in timestamp format.
	LastModifiedDate *time.Time

	// A comment or description about the repository.
	RepositoryDescription *string

	// The ID of the repository.
	RepositoryId *string

	// The repository's name.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Information about a repository name and ID.
type RepositoryNameIdPair struct {

	// The ID associated with the repository.
	RepositoryId *string

	// The name associated with the repository.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Information about a trigger for a repository.
//
// If you want to receive notifications about repository events, consider using
// notifications instead of triggers. For more information, see [Configuring notifications for repository events].
//
// [Configuring notifications for repository events]: https://docs.aws.amazon.com/codecommit/latest/userguide/how-to-repository-email.html
type RepositoryTrigger struct {

	// The ARN of the resource that is the target for a trigger (for example, the ARN
	// of a topic in Amazon SNS).
	//
	// This member is required.
	DestinationArn *string

	// The repository events that cause the trigger to run actions in another service,
	// such as sending a notification through Amazon SNS.
	//
	// The valid value "all" cannot be used with any other values.
	//
	// This member is required.
	Events []RepositoryTriggerEventEnum

	// The name of the trigger.
	//
	// This member is required.
	Name *string

	// The branches to be included in the trigger configuration. If you specify an
	// empty array, the trigger applies to all branches.
	//
	// Although no content is required in the array, you must include the array itself.
	Branches []string

	// Any custom data associated with the trigger to be included in the information
	// sent to the target of the trigger.
	CustomData *string

	noSmithyDocumentSerde
}

// A trigger failed to run.
type RepositoryTriggerExecutionFailure struct {

	// Message information about the trigger that did not run.
	FailureMessage *string

	// The name of the trigger that did not run.
	Trigger *string

	noSmithyDocumentSerde
}

// Information about the file mode changes.
type SetFileModeEntry struct {

	// The file mode for the file.
	//
	// This member is required.
	FileMode FileModeTypeEnum

	// The full path to the file, including the name of the file.
	//
	// This member is required.
	FilePath *string

	noSmithyDocumentSerde
}

// Information about a source file that is part of changes made in a commit.
type SourceFileSpecifier struct {

	// The full path to the file, including the name of the file.
	//
	// This member is required.
	FilePath *string

	// Whether to remove the source file from the parent commit.
	IsMove bool

	noSmithyDocumentSerde
}

// Returns information about a submodule reference in a repository folder.
type SubModule struct {

	// The fully qualified path to the folder that contains the reference to the
	// submodule.
	AbsolutePath *string

	// The commit ID that contains the reference to the submodule.
	CommitId *string

	// The relative path of the submodule from the folder where the query originated.
	RelativePath *string

	noSmithyDocumentSerde
}

// Returns information about a symbolic link in a repository folder.
type SymbolicLink struct {

	// The fully qualified path to the folder that contains the symbolic link.
	AbsolutePath *string

	// The blob ID that contains the information about the symbolic link.
	BlobId *string

	// The file mode permissions of the blob that cotains information about the
	// symbolic link.
	FileMode FileModeTypeEnum

	// The relative path of the symbolic link from the folder where the query
	// originated.
	RelativePath *string

	noSmithyDocumentSerde
}

// Returns information about a target for a pull request.
type Target struct {

	// The name of the repository that contains the pull request.
	//
	// This member is required.
	RepositoryName *string

	// The branch of the repository that contains the changes for the pull request.
	// Also known as the source branch.
	//
	// This member is required.
	SourceReference *string

	// The branch of the repository where the pull request changes are merged. Also
	// known as the destination branch.
	DestinationReference *string

	noSmithyDocumentSerde
}

// Information about the user who made a specified commit.
type UserInfo struct {

	// The date when the specified commit was commited, in timestamp format with GMT
	// offset.
	Date *string

	// The email address associated with the user who made the commit, if any.
	Email *string

	// The name of the user who made the specified commit.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
