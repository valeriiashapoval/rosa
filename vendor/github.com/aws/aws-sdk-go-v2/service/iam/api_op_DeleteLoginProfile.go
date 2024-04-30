// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the password for the specified IAM user, For more information, see
// Managing passwords for IAM users (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_admin-change-user.html)
// . You can use the CLI, the Amazon Web Services API, or the Users page in the IAM
// console to delete a password for any IAM user. You can use ChangePassword to
// update, but not delete, your own password in the My Security Credentials page in
// the Amazon Web Services Management Console. Deleting a user's password does not
// prevent a user from accessing Amazon Web Services through the command line
// interface or the API. To prevent all user access, you must also either make any
// access keys inactive or delete them. For more information about making keys
// inactive or deleting them, see UpdateAccessKey and DeleteAccessKey .
func (c *Client) DeleteLoginProfile(ctx context.Context, params *DeleteLoginProfileInput, optFns ...func(*Options)) (*DeleteLoginProfileOutput, error) {
	if params == nil {
		params = &DeleteLoginProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLoginProfile", params, optFns, c.addOperationDeleteLoginProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLoginProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLoginProfileInput struct {

	// The name of the user whose password you want to delete. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex) ) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

type DeleteLoginProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteLoginProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteLoginProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteLoginProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteLoginProfile"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteLoginProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLoginProfile(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteLoginProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteLoginProfile",
	}
}
