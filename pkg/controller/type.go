package controller

import (
	"time"

	rapi "github.com/appscode/k8s-addons/api"
	tcs "github.com/appscode/k8s-addons/client/clientset"
	"gopkg.in/robfig/cron.v2"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/client/record"
)

const (
	BackupConfig          = "backup.appscode.com/config"
	ContainerName         = "restic-sidecar"
	RestikNamespace       = "RESTIK_NAMESPACE"
	RestikResourceName    = "RESTIK_RESOURCE_NAME"
	RESTIC_PASSWORD       = "RESTIC_PASSWORD"
	ReplicationController = "ReplicationController"
	ReplicaSet            = "ReplicaSet"
	Deployment            = "Deployment"
	DaemonSet             = "DaemonSet"
	StatefulSet           = "StatefulSet"
	Password              = "password"
	ImageAnnotation       = "backup.appscode.com/image"
	Force                 = "force"
)

const (
	EventReasonInvalidCronExpression         = "InvalidCronExpression"
	EventReasonSuccessfulCronExpressionReset = "SuccessfulCronExpressionReset"
	EventReasonSuccessfulBackup              = "SuccessfulBackup"
	EventReasonFailedToBackup                = "FailedBackup"
	EventReasonFailedToRetention             = "FailedRetention"
	EventReasonFailedToUpdate                = "FailedUpdateBackup"
	EventReasonFailedCronJob                 = "FailedCronJob"
)

type Controller struct {
	ExtClient tcs.AppsCodeExtensionInterface
	Client    clientset.Interface
	// sync time to sync the list.
	SyncPeriod time.Duration
	// image of sidecar container
	Image string
}

type cronController struct {
	extClient     tcs.AppsCodeExtensionInterface
	kubeClient    clientset.Interface
	tprName       string
	namespace     string
	crons         *cron.Cron
	backup        *rapi.Backup
	eventRecorder record.EventRecorder
}