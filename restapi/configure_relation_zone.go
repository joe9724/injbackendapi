// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"injbackendapi/restapi/operations"
	"injbackendapi/restapi/operations/comment"
	"injbackendapi/restapi/operations/file_upload"
	"injbackendapi/restapi/operations/status"
	"injbackendapi/restapi/operations/topic"
	"injbackendapi/restapi/operations/user"
	"injbackendapi/restapi/operations/zone"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.json

func configureFlags(api *operations.RelationZoneAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RelationZoneAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.CommentNrCommentsCreateHandler = comment.NrCommentsCreateHandlerFunc(func(params comment.NrCommentsCreateParams) middleware.Responder {
		return middleware.NotImplemented("operation comment.NrCommentsCreate has not yet been implemented")
	})
	api.CommentNrCommentsDeleteHandler = comment.NrCommentsDeleteHandlerFunc(func(params comment.NrCommentsDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation comment.NrCommentsDelete has not yet been implemented")
	})
	api.CommentNrCommentsListHandler = comment.NrCommentsListHandlerFunc(func(params comment.NrCommentsListParams) middleware.Responder {
		return middleware.NotImplemented("operation comment.NrCommentsList has not yet been implemented")
	})
	api.CommentNrCommentsReplyHandler = comment.NrCommentsReplyHandlerFunc(func(params comment.NrCommentsReplyParams) middleware.Responder {
		return middleware.NotImplemented("operation comment.NrCommentsReply has not yet been implemented")
	})
	api.FileUploadNrFileUploadHandler = file_upload.NrFileUploadHandlerFunc(func(params file_upload.NrFileUploadParams) middleware.Responder {
		return middleware.NotImplemented("operation file_upload.NrFileUpload has not yet been implemented")
	})
	api.StatusNrStatusCreateHandler = status.NrStatusCreateHandlerFunc(func(params status.NrStatusCreateParams) middleware.Responder {
		return middleware.NotImplemented("operation status.NrStatusCreate has not yet been implemented")
	})
	api.StatusNrStatusLikeHandler = status.NrStatusLikeHandlerFunc(func(params status.NrStatusLikeParams) middleware.Responder {
		return middleware.NotImplemented("operation status.NrStatusLike has not yet been implemented")
	})
	api.StatusNrStatusRetweetHandler = status.NrStatusRetweetHandlerFunc(func(params status.NrStatusRetweetParams) middleware.Responder {
		return middleware.NotImplemented("operation status.NrStatusRetweet has not yet been implemented")
	})
	api.StatusNrStatusSearchHandler = status.NrStatusSearchHandlerFunc(func(params status.NrStatusSearchParams) middleware.Responder {
		return middleware.NotImplemented("operation status.NrStatusSearch has not yet been implemented")
	})
	api.StatusNrStatusShowHandler = status.NrStatusShowHandlerFunc(func(params status.NrStatusShowParams) middleware.Responder {
		return middleware.NotImplemented("operation status.NrStatusShow has not yet been implemented")
	})
	api.StatusNrStatusTagsHandler = status.NrStatusTagsHandlerFunc(func(params status.NrStatusTagsParams) middleware.Responder {
		return middleware.NotImplemented("operation status.NrStatusTags has not yet been implemented")
	})
	api.StatusNrStatusTimelineHandler = status.NrStatusTimelineHandlerFunc(func(params status.NrStatusTimelineParams) middleware.Responder {
		return middleware.NotImplemented("operation status.NrStatusTimeline has not yet been implemented")
	})
	api.StatusNrStatusUnlikeHandler = status.NrStatusUnlikeHandlerFunc(func(params status.NrStatusUnlikeParams) middleware.Responder {
		return middleware.NotImplemented("operation status.NrStatusUnlike has not yet been implemented")
	})
	api.TopicNrTopicCreateHandler = topic.NrTopicCreateHandlerFunc(func(params topic.NrTopicCreateParams) middleware.Responder {
		return middleware.NotImplemented("operation topic.NrTopicCreate has not yet been implemented")
	})
	api.TopicNrTopicFollowHandler = topic.NrTopicFollowHandlerFunc(func(params topic.NrTopicFollowParams) middleware.Responder {
		return middleware.NotImplemented("operation topic.NrTopicFollow has not yet been implemented")
	})
	api.TopicNrTopicListHandler = topic.NrTopicListHandlerFunc(func(params topic.NrTopicListParams) middleware.Responder {
		return middleware.NotImplemented("operation topic.NrTopicList has not yet been implemented")
	})
	api.TopicNrTopicMatchHandler = topic.NrTopicMatchHandlerFunc(func(params topic.NrTopicMatchParams) middleware.Responder {
		return middleware.NotImplemented("operation topic.NrTopicMatch has not yet been implemented")
	})
	api.TopicNrTopicSearchHandler = topic.NrTopicSearchHandlerFunc(func(params topic.NrTopicSearchParams) middleware.Responder {
		return middleware.NotImplemented("operation topic.NrTopicSearch has not yet been implemented")
	})
	api.TopicNrTopicShowHandler = topic.NrTopicShowHandlerFunc(func(params topic.NrTopicShowParams) middleware.Responder {
		return middleware.NotImplemented("operation topic.NrTopicShow has not yet been implemented")
	})
	api.TopicNrTopicStatusHandler = topic.NrTopicStatusHandlerFunc(func(params topic.NrTopicStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation topic.NrTopicStatus has not yet been implemented")
	})
	api.TopicNrTopicUnfollowHandler = topic.NrTopicUnfollowHandlerFunc(func(params topic.NrTopicUnfollowParams) middleware.Responder {
		return middleware.NotImplemented("operation topic.NrTopicUnfollow has not yet been implemented")
	})
	api.UserNrUserAggreeToFriendHandler = user.NrUserAggreeToFriendHandlerFunc(func(params user.NrUserAggreeToFriendParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserAggreeToFriend has not yet been implemented")
	})
	api.UserNrUserApplyFriendHandler = user.NrUserApplyFriendHandlerFunc(func(params user.NrUserApplyFriendParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserApplyFriend has not yet been implemented")
	})
	api.UserNrUserDeleteFriendHandler = user.NrUserDeleteFriendHandlerFunc(func(params user.NrUserDeleteFriendParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserDeleteFriend has not yet been implemented")
	})
	api.UserNrUserFansHandler = user.NrUserFansHandlerFunc(func(params user.NrUserFansParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserFans has not yet been implemented")
	})
	api.UserNrUserFollowHandler = user.NrUserFollowHandlerFunc(func(params user.NrUserFollowParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserFollow has not yet been implemented")
	})
	api.UserNrUserFollowsHandler = user.NrUserFollowsHandlerFunc(func(params user.NrUserFollowsParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserFollows has not yet been implemented")
	})
	api.UserNrUserFriendsHandler = user.NrUserFriendsHandlerFunc(func(params user.NrUserFriendsParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserFriends has not yet been implemented")
	})
	api.UserNrUserPhotosHandler = user.NrUserPhotosHandlerFunc(func(params user.NrUserPhotosParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserPhotos has not yet been implemented")
	})
	api.UserNrUserRefuseToFriendHandler = user.NrUserRefuseToFriendHandlerFunc(func(params user.NrUserRefuseToFriendParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserRefuseToFriend has not yet been implemented")
	})
	api.UserNrUserShowHandler = user.NrUserShowHandlerFunc(func(params user.NrUserShowParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserShow has not yet been implemented")
	})
	api.UserNrUserStatusHandler = user.NrUserStatusHandlerFunc(func(params user.NrUserStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserStatus has not yet been implemented")
	})
	api.UserNrUserUnfollowHandler = user.NrUserUnfollowHandlerFunc(func(params user.NrUserUnfollowParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserUnfollow has not yet been implemented")
	})
	api.UserNrUserZonesHandler = user.NrUserZonesHandlerFunc(func(params user.NrUserZonesParams) middleware.Responder {
		return middleware.NotImplemented("operation user.NrUserZones has not yet been implemented")
	})
	api.ZoneNrZoneAggreeHandler = zone.NrZoneAggreeHandlerFunc(func(params zone.NrZoneAggreeParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneAggree has not yet been implemented")
	})
	api.ZoneNrZoneApplyHandler = zone.NrZoneApplyHandlerFunc(func(params zone.NrZoneApplyParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneApply has not yet been implemented")
	})
	api.ZoneNrZoneCreateHandler = zone.NrZoneCreateHandlerFunc(func(params zone.NrZoneCreateParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneCreate has not yet been implemented")
	})
	api.ZoneNrZoneDropHandler = zone.NrZoneDropHandlerFunc(func(params zone.NrZoneDropParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneDrop has not yet been implemented")
	})
	api.ZoneNrZoneInviteHandler = zone.NrZoneInviteHandlerFunc(func(params zone.NrZoneInviteParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneInvite has not yet been implemented")
	})
	api.ZoneNrZoneListHandler = zone.NrZoneListHandlerFunc(func(params zone.NrZoneListParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneList has not yet been implemented")
	})
	api.ZoneNrZoneMembersHandler = zone.NrZoneMembersHandlerFunc(func(params zone.NrZoneMembersParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneMembers has not yet been implemented")
	})
	api.ZoneNrZoneQuiteHandler = zone.NrZoneQuiteHandlerFunc(func(params zone.NrZoneQuiteParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneQuite has not yet been implemented")
	})
	api.ZoneNrZoneRecommendHandler = zone.NrZoneRecommendHandlerFunc(func(params zone.NrZoneRecommendParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneRecommend has not yet been implemented")
	})
	api.ZoneNrZoneRefuseHandler = zone.NrZoneRefuseHandlerFunc(func(params zone.NrZoneRefuseParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneRefuse has not yet been implemented")
	})
	api.ZoneNrZoneReportHandler = zone.NrZoneReportHandlerFunc(func(params zone.NrZoneReportParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneReport has not yet been implemented")
	})
	api.ZoneNrZoneShowHandler = zone.NrZoneShowHandlerFunc(func(params zone.NrZoneShowParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneShow has not yet been implemented")
	})
	api.ZoneNrZoneStatusHandler = zone.NrZoneStatusHandlerFunc(func(params zone.NrZoneStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneStatus has not yet been implemented")
	})
	api.ZoneNrZoneUpdateHandler = zone.NrZoneUpdateHandlerFunc(func(params zone.NrZoneUpdateParams) middleware.Responder {
		return middleware.NotImplemented("operation zone.NrZoneUpdate has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
