
package docker
import "docker.io/go-docker"

Package docker is the official Go client for the Docker API.

For more information about the Docker API, see the documentation: https://docs.docker.com/develop/api

Usage
You use the library by creating a client object and calling methods on it. The client can be created either from environment variables with NewEnvClient, or configured manually with NewClient.

For example, to list running containers (the equivalent of "docker ps"):

package main

import (
	"context"
	"fmt"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

func main() {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
Index
Constants
Variables
func CheckRedirect(req *http.Request, via []*http.Request) error
func ErrorConnectionFailed(host string) error
func IsErrConfigNotFound(err error) bool
func IsErrConnectionFailed(err error) bool
func IsErrContainerNotFound(err error) bool
func IsErrImageNotFound(err error) bool
func IsErrNetworkNotFound(err error) bool
func IsErrNodeNotFound(err error) bool
func IsErrNotFound(err error) bool
func IsErrNotImplemented(err error) bool
func IsErrPluginNotFound(err error) bool
func IsErrPluginPermissionDenied(err error) bool
func IsErrSecretNotFound(err error) bool
func IsErrServiceNotFound(err error) bool
func IsErrTaskNotFound(err error) bool
func IsErrUnauthorized(err error) bool
func IsErrVolumeNotFound(err error) bool
func ParseHost(host string) (string, string, string, error)
func ParseHostURL(host string) (*url.URL, error)
type APIClient
type CheckpointAPIClient
type Client
func NewClient(host string, version string, client *http.Client, httpHeaders map[string]string) (*Client, error)
func NewEnvClient() (*Client, error)
func (cli *Client) BuildCachePrune(ctx context.Context) (*types.BuildCachePruneReport, error)
func (cli *Client) CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) error
func (cli *Client) CheckpointDelete(ctx context.Context, containerID string, options types.CheckpointDeleteOptions) error
func (cli *Client) CheckpointList(ctx context.Context, container string, options types.CheckpointListOptions) ([]types.Checkpoint, error)
func (cli *Client) ClientVersion() string
func (cli *Client) Close() error
func (cli *Client) ConfigCreate(ctx context.Context, config swarm.ConfigSpec) (types.ConfigCreateResponse, error)
func (cli *Client) ConfigInspectWithRaw(ctx context.Context, id string) (swarm.Config, []byte, error)
func (cli *Client) ConfigList(ctx context.Context, options types.ConfigListOptions) ([]swarm.Config, error)
func (cli *Client) ConfigRemove(ctx context.Context, id string) error
func (cli *Client) ConfigUpdate(ctx context.Context, id string, version swarm.Version, config swarm.ConfigSpec) error
func (cli *Client) ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
func (cli *Client) ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
func (cli *Client) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
func (cli *Client) ContainerDiff(ctx context.Context, containerID string) ([]container.ContainerChangeResponseItem, error)
func (cli *Client) ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error)
func (cli *Client) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
func (cli *Client) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
func (cli *Client) ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
func (cli *Client) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
func (cli *Client) ContainerExport(ctx context.Context, containerID string) (io.ReadCloser, error)
func (cli *Client) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
func (cli *Client) ContainerInspectWithRaw(ctx context.Context, containerID string, getSize bool) (types.ContainerJSON, []byte, error)
func (cli *Client) ContainerKill(ctx context.Context, containerID, signal string) error
func (cli *Client) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
func (cli *Client) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
func (cli *Client) ContainerPause(ctx context.Context, containerID string) error
func (cli *Client) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
func (cli *Client) ContainerRename(ctx context.Context, containerID, newContainerName string) error
func (cli *Client) ContainerResize(ctx context.Context, containerID string, options types.ResizeOptions) error
func (cli *Client) ContainerRestart(ctx context.Context, containerID string, timeout *time.Duration) error
func (cli *Client) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error
func (cli *Client) ContainerStatPath(ctx context.Context, containerID, path string) (types.ContainerPathStat, error)
func (cli *Client) ContainerStats(ctx context.Context, containerID string, stream bool) (types.ContainerStats, error)
func (cli *Client) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
func (cli *Client) ContainerTop(ctx context.Context, containerID string, arguments []string) (container.ContainerTopOKBody, error)
func (cli *Client) ContainerUnpause(ctx context.Context, containerID string) error
func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error)
func (cli *Client) ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error)
func (cli *Client) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
func (cli *Client) CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error)
func (cli *Client) CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error
func (cli *Client) CustomHTTPHeaders() map[string]string
func (cli *Client) DaemonHost() string
func (cli *Client) DialSession(ctx context.Context, proto string, meta map[string][]string) (net.Conn, error)
func (cli *Client) DiskUsage(ctx context.Context) (types.DiskUsage, error)
func (cli *Client) DistributionInspect(ctx context.Context, image, encodedRegistryAuth string) (registrytypes.DistributionInspect, error)
func (cli *Client) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
func (cli *Client) ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
func (cli *Client) ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error)
func (cli *Client) ImageHistory(ctx context.Context, imageID string) ([]image.HistoryResponseItem, error)
func (cli *Client) ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error)
func (cli *Client) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error)
func (cli *Client) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
func (cli *Client) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)
func (cli *Client) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)
func (cli *Client) ImagePush(ctx context.Context, image string, options types.ImagePushOptions) (io.ReadCloser, error)
func (cli *Client) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)
func (cli *Client) ImageSave(ctx context.Context, imageIDs []string) (io.ReadCloser, error)
func (cli *Client) ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error)
func (cli *Client) ImageTag(ctx context.Context, source, target string) error
func (cli *Client) ImagesPrune(ctx context.Context, pruneFilters filters.Args) (types.ImagesPruneReport, error)
func (cli *Client) Info(ctx context.Context) (types.Info, error)
func (cli *Client) NegotiateAPIVersion(ctx context.Context)
func (cli *Client) NegotiateAPIVersionPing(p types.Ping)
func (cli *Client) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error
func (cli *Client) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)
func (cli *Client) NetworkDisconnect(ctx context.Context, networkID, containerID string, force bool) error
func (cli *Client) NetworkInspect(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, error)
func (cli *Client) NetworkInspectWithRaw(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error)
func (cli *Client) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error)
func (cli *Client) NetworkRemove(ctx context.Context, networkID string) error
func (cli *Client) NetworksPrune(ctx context.Context, pruneFilters filters.Args) (types.NetworksPruneReport, error)
func (cli *Client) NewVersionError(APIrequired, feature string) error
func (cli *Client) NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error)
func (cli *Client) NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error)
func (cli *Client) NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error
func (cli *Client) NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error
func (cli *Client) Ping(ctx context.Context) (types.Ping, error)
func (cli *Client) PluginCreate(ctx context.Context, createContext io.Reader, createOptions types.PluginCreateOptions) error
func (cli *Client) PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error
func (cli *Client) PluginEnable(ctx context.Context, name string, options types.PluginEnableOptions) error
func (cli *Client) PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error)
func (cli *Client) PluginInstall(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
func (cli *Client) PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error)
func (cli *Client) PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error)
func (cli *Client) PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error
func (cli *Client) PluginSet(ctx context.Context, name string, args []string) error
func (cli *Client) PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
func (cli *Client) RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)
func (cli *Client) SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error)
func (cli *Client) SecretInspectWithRaw(ctx context.Context, id string) (swarm.Secret, []byte, error)
func (cli *Client) SecretList(ctx context.Context, options types.SecretListOptions) ([]swarm.Secret, error)
func (cli *Client) SecretRemove(ctx context.Context, id string) error
func (cli *Client) SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error
func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error)
func (cli *Client) ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error)
func (cli *Client) ServiceInspectWithRaw(ctx context.Context, serviceID string, opts types.ServiceInspectOptions) (swarm.Service, []byte, error)
func (cli *Client) ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
func (cli *Client) ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
func (cli *Client) ServiceRemove(ctx context.Context, serviceID string) error
func (cli *Client) ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) (types.ServiceUpdateResponse, error)
func (cli *Client) SetCustomHTTPHeaders(headers map[string]string)
func (cli *Client) SwarmGetUnlockKey(ctx context.Context) (types.SwarmUnlockKeyResponse, error)
func (cli *Client) SwarmInit(ctx context.Context, req swarm.InitRequest) (string, error)
func (cli *Client) SwarmInspect(ctx context.Context) (swarm.Swarm, error)
func (cli *Client) SwarmJoin(ctx context.Context, req swarm.JoinRequest) error
func (cli *Client) SwarmLeave(ctx context.Context, force bool) error
func (cli *Client) SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error
func (cli *Client) SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error
func (cli *Client) TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error)
func (cli *Client) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error)
func (cli *Client) TaskLogs(ctx context.Context, taskID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
func (cli *Client) VolumeCreate(ctx context.Context, options volumetypes.VolumesCreateBody) (types.Volume, error)
func (cli *Client) VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error)
func (cli *Client) VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error)
func (cli *Client) VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumesListOKBody, error)
func (cli *Client) VolumeRemove(ctx context.Context, volumeID string, force bool) error
func (cli *Client) VolumesPrune(ctx context.Context, pruneFilters filters.Args) (types.VolumesPruneReport, error)
type CommonAPIClient
type ConfigAPIClient
type ContainerAPIClient
type DistributionAPIClient
type ImageAPIClient
type NetworkAPIClient
type NodeAPIClient
type PluginAPIClient
type SecretAPIClient
type ServiceAPIClient
type SwarmAPIClient
type SystemAPIClient
type VolumeAPIClient
Examples
Client.ContainerLogs (WithTimeout)
Client.ContainerWait (WithTimeout)
Client.ServiceLogs (WithTimeout)
Package Files
build_prune.go checkpoint_create.go checkpoint_delete.go checkpoint_list.go client.go client_unix.go config_create.go config_inspect.go config_list.go config_remove.go config_update.go container_attach.go container_commit.go container_copy.go container_create.go container_diff.go container_exec.go container_export.go container_inspect.go container_kill.go container_list.go container_logs.go container_pause.go container_prune.go container_remove.go container_rename.go container_resize.go container_restart.go container_start.go container_stats.go container_stop.go container_top.go container_unpause.go container_update.go container_wait.go disk_usage.go distribution_inspect.go doc.go errors.go events.go hijack.go image_build.go image_create.go image_history.go image_import.go image_inspect.go image_list.go image_load.go image_prune.go image_pull.go image_push.go image_remove.go image_save.go image_search.go image_tag.go info.go interface.go interface_experimental.go interface_stable.go login.go network_connect.go network_create.go network_disconnect.go network_inspect.go network_list.go network_prune.go network_remove.go node_inspect.go node_list.go node_remove.go node_update.go ping.go plugin_create.go plugin_disable.go plugin_enable.go plugin_inspect.go plugin_install.go plugin_list.go plugin_push.go plugin_remove.go plugin_set.go plugin_upgrade.go request.go secret_create.go secret_inspect.go secret_list.go secret_remove.go secret_update.go service_create.go service_inspect.go service_list.go service_logs.go service_remove.go service_update.go session.go swarm_get_unlock_key.go swarm_init.go swarm_inspect.go swarm_join.go swarm_leave.go swarm_unlock.go swarm_update.go task_inspect.go task_list.go task_logs.go tlsconfig_clone.go transport.go utils.go version.go volume_create.go volume_inspect.go volume_list.go volume_prune.go volume_remove.go

Constants
const DefaultDockerHost = "unix:///var/run/docker.sock"
DefaultDockerHost defines os specific default if DOCKER_HOST is unset

Variables
var ErrRedirect = errors.New("unexpected redirect in response")
ErrRedirect is the error returned by checkRedirect when the request is non-GET.

func CheckRedirect
func CheckRedirect(req *http.Request, via []*http.Request) error
CheckRedirect specifies the policy for dealing with redirect responses: If the request is non-GET return `ErrRedirect`. Otherwise use the last response.

Go 1.8 changes behavior for HTTP redirects (specifically 301, 307, and 308) in the client . The Docker client (and by extension docker API client) can be made to to send a request like POST /containers//start where what would normally be in the name section of the URL is empty. This triggers an HTTP 301 from the daemon. In go 1.8 this 301 will be converted to a GET request, and ends up getting a 404 from the daemon. This behavior change manifests in the client in that before the 301 was not followed and the client did not generate an error, but now results in a message like Error response from daemon: page not found.

func ErrorConnectionFailed
func ErrorConnectionFailed(host string) error
ErrorConnectionFailed returns an error with host in the error message when connection to docker daemon failed.

func IsErrConfigNotFound
func IsErrConfigNotFound(err error) bool
IsErrConfigNotFound returns true if the error is caused when a config is not found.

Deprecated: Use IsErrNotFound

func IsErrConnectionFailed
func IsErrConnectionFailed(err error) bool
IsErrConnectionFailed returns true if the error is caused by connection failed.

func IsErrContainerNotFound
func IsErrContainerNotFound(err error) bool
IsErrContainerNotFound returns true if the error is caused when a container is not found in the docker host.

Deprecated: Use IsErrNotFound

func IsErrImageNotFound
func IsErrImageNotFound(err error) bool
IsErrImageNotFound returns true if the error is caused when an image is not found in the docker host.

Deprecated: Use IsErrNotFound

func IsErrNetworkNotFound
func IsErrNetworkNotFound(err error) bool
IsErrNetworkNotFound returns true if the error is caused when a network is not found in the docker host.

Deprecated: Use IsErrNotFound

func IsErrNodeNotFound
func IsErrNodeNotFound(err error) bool
IsErrNodeNotFound returns true if the error is caused when a node is not found.

Deprecated: Use IsErrNotFound

func IsErrNotFound
func IsErrNotFound(err error) bool
IsErrNotFound returns true if the error is a NotFound error, which is returned by the API when some object is not found.

func IsErrNotImplemented
func IsErrNotImplemented(err error) bool
IsErrNotImplemented returns true if the error is a NotImplemented error. This is returned by the API when a requested feature has not been implemented.

func IsErrPluginNotFound
func IsErrPluginNotFound(err error) bool
IsErrPluginNotFound returns true if the error is caused when a plugin is not found in the docker host.

Deprecated: Use IsErrNotFound

func IsErrPluginPermissionDenied
func IsErrPluginPermissionDenied(err error) bool
IsErrPluginPermissionDenied returns true if the error is caused when a user denies a plugin's permissions

func IsErrSecretNotFound
func IsErrSecretNotFound(err error) bool
IsErrSecretNotFound returns true if the error is caused when a secret is not found.

Deprecated: Use IsErrNotFound

func IsErrServiceNotFound
func IsErrServiceNotFound(err error) bool
IsErrServiceNotFound returns true if the error is caused when a service is not found.

Deprecated: Use IsErrNotFound

func IsErrTaskNotFound
func IsErrTaskNotFound(err error) bool
IsErrTaskNotFound returns true if the error is caused when a task is not found.

Deprecated: Use IsErrNotFound

func IsErrUnauthorized
func IsErrUnauthorized(err error) bool
IsErrUnauthorized returns true if the error is caused when a remote registry authentication fails

func IsErrVolumeNotFound
func IsErrVolumeNotFound(err error) bool
IsErrVolumeNotFound returns true if the error is caused when a volume is not found in the docker host.

Deprecated: Use IsErrNotFound

func ParseHost
func ParseHost(host string) (string, string, string, error)
ParseHost parses a url string, validates the strings is a host url, and returns the parsed host as: protocol, address, and base path Deprecated: use ParseHostURL

func ParseHostURL
func ParseHostURL(host string) (*url.URL, error)
ParseHostURL parses a url string, validates the string is a host url, and returns the parsed URL

type APIClient
type APIClient interface {
    CommonAPIClient
    // contains filtered or unexported methods
}
APIClient is an interface that clients that talk with a docker server must implement.

type CheckpointAPIClient
type CheckpointAPIClient interface {
    CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) error
    CheckpointDelete(ctx context.Context, container string, options types.CheckpointDeleteOptions) error
    CheckpointList(ctx context.Context, container string, options types.CheckpointListOptions) ([]types.Checkpoint, error)
}
CheckpointAPIClient defines API client methods for the checkpoints

type Client
type Client struct {
    // contains filtered or unexported fields
}
Client is the API client that performs all operations against a docker server.

func NewClient
func NewClient(host string, version string, client *http.Client, httpHeaders map[string]string) (*Client, error)
NewClient initializes a new API client for the given host and API version. It uses the given http client as transport. It also initializes the custom http headers to add to each request.

It won't send any version information if the version number is empty. It is highly recommended that you set a version or your client may break if the server is upgraded.

func NewEnvClient
func NewEnvClient() (*Client, error)
NewEnvClient initializes a new API client based on environment variables. Use DOCKER_HOST to set the url to the docker server. Use DOCKER_API_VERSION to set the version of the API to reach, leave empty for latest. Use DOCKER_CERT_PATH to load the TLS certificates from. Use DOCKER_TLS_VERIFY to enable or disable TLS verification, off by default.

func (*Client) BuildCachePrune
func (cli *Client) BuildCachePrune(ctx context.Context) (*types.BuildCachePruneReport, error)
BuildCachePrune requests the daemon to delete unused cache data

func (*Client) CheckpointCreate
func (cli *Client) CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) error
CheckpointCreate creates a checkpoint from the given container with the given name

func (*Client) CheckpointDelete
func (cli *Client) CheckpointDelete(ctx context.Context, containerID string, options types.CheckpointDeleteOptions) error
CheckpointDelete deletes the checkpoint with the given name from the given container

func (*Client) CheckpointList
func (cli *Client) CheckpointList(ctx context.Context, container string, options types.CheckpointListOptions) ([]types.Checkpoint, error)
CheckpointList returns the checkpoints of the given container in the docker host

func (*Client) ClientVersion
func (cli *Client) ClientVersion() string
ClientVersion returns the API version used by this client.

func (*Client) Close
func (cli *Client) Close() error
Close the transport used by the client

func (*Client) ConfigCreate
func (cli *Client) ConfigCreate(ctx context.Context, config swarm.ConfigSpec) (types.ConfigCreateResponse, error)
ConfigCreate creates a new Config.

func (*Client) ConfigInspectWithRaw
func (cli *Client) ConfigInspectWithRaw(ctx context.Context, id string) (swarm.Config, []byte, error)
ConfigInspectWithRaw returns the config information with raw data

func (*Client) ConfigList
func (cli *Client) ConfigList(ctx context.Context, options types.ConfigListOptions) ([]swarm.Config, error)
ConfigList returns the list of configs.

func (*Client) ConfigRemove
func (cli *Client) ConfigRemove(ctx context.Context, id string) error
ConfigRemove removes a Config.

func (*Client) ConfigUpdate
func (cli *Client) ConfigUpdate(ctx context.Context, id string, version swarm.Version, config swarm.ConfigSpec) error
ConfigUpdate attempts to update a Config

func (*Client) ContainerAttach
func (cli *Client) ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
ContainerAttach attaches a connection to a container in the server. It returns a types.HijackedConnection with the hijacked connection and the a reader to get output. It's up to the called to close the hijacked connection by calling types.HijackedResponse.Close.

The stream format on the response will be in one of two formats:

If the container is using a TTY, there is only a single stream (stdout), and data is copied directly from the container output stream, no extra multiplexing or headers.

If the container is *not* using a TTY, streams for stdout and stderr are multiplexed. The format of the multiplexed stream is as follows:

[8]byte{STREAM_TYPE, 0, 0, 0, SIZE1, SIZE2, SIZE3, SIZE4}[]byte{OUTPUT}
STREAM_TYPE can be 1 for stdout and 2 for stderr

SIZE1, SIZE2, SIZE3, and SIZE4 are four bytes of uint32 encoded as big endian. This is the size of OUTPUT.

You can use github.com/docker/docker/pkg/stdcopy.StdCopy to demultiplex this stream.

func (*Client) ContainerCommit
func (cli *Client) ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
ContainerCommit applies changes into a container and creates a new tagged image.

func (*Client) ContainerCreate
func (cli *Client) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
ContainerCreate creates a new container based in the given configuration. It can be associated with a name, but it's not mandatory.

func (*Client) ContainerDiff
func (cli *Client) ContainerDiff(ctx context.Context, containerID string) ([]container.ContainerChangeResponseItem, error)
ContainerDiff shows differences in a container filesystem since it was started.

func (*Client) ContainerExecAttach
func (cli *Client) ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error)
ContainerExecAttach attaches a connection to an exec process in the server. It returns a types.HijackedConnection with the hijacked connection and the a reader to get output. It's up to the called to close the hijacked connection by calling types.HijackedResponse.Close.

func (*Client) ContainerExecCreate
func (cli *Client) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
ContainerExecCreate creates a new exec configuration to run an exec process.

func (*Client) ContainerExecInspect
func (cli *Client) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
ContainerExecInspect returns information about a specific exec process on the docker host.

func (*Client) ContainerExecResize
func (cli *Client) ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
ContainerExecResize changes the size of the tty for an exec process running inside a container.

func (*Client) ContainerExecStart
func (cli *Client) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
ContainerExecStart starts an exec process already created in the docker host.

func (*Client) ContainerExport
func (cli *Client) ContainerExport(ctx context.Context, containerID string) (io.ReadCloser, error)
ContainerExport retrieves the raw contents of a container and returns them as an io.ReadCloser. It's up to the caller to close the stream.

func (*Client) ContainerInspect
func (cli *Client) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
ContainerInspect returns the container information.

func (*Client) ContainerInspectWithRaw
func (cli *Client) ContainerInspectWithRaw(ctx context.Context, containerID string, getSize bool) (types.ContainerJSON, []byte, error)
ContainerInspectWithRaw returns the container information and its raw representation.

func (*Client) ContainerKill
func (cli *Client) ContainerKill(ctx context.Context, containerID, signal string) error
ContainerKill terminates the container process but does not remove the container from the docker host.

func (*Client) ContainerList
func (cli *Client) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
ContainerList returns the list of containers in the docker host.

func (*Client) ContainerLogs
func (cli *Client) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
ContainerLogs returns the logs generated by a container in an io.ReadCloser. It's up to the caller to close the stream.

The stream format on the response will be in one of two formats:

If the container is using a TTY, there is only a single stream (stdout), and data is copied directly from the container output stream, no extra multiplexing or headers.

If the container is *not* using a TTY, streams for stdout and stderr are multiplexed. The format of the multiplexed stream is as follows:

[8]byte{STREAM_TYPE, 0, 0, 0, SIZE1, SIZE2, SIZE3, SIZE4}[]byte{OUTPUT}
STREAM_TYPE can be 1 for stdout and 2 for stderr

SIZE1, SIZE2, SIZE3, and SIZE4 are four bytes of uint32 encoded as big endian. This is the size of OUTPUT.

You can use github.com/docker/docker/pkg/stdcopy.StdCopy to demultiplex this stream.

Example (WithTimeout)
func (*Client) ContainerPause
func (cli *Client) ContainerPause(ctx context.Context, containerID string) error
ContainerPause pauses the main process of a given container without terminating it.

func (*Client) ContainerRemove
func (cli *Client) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
ContainerRemove kills and removes a container from the docker host.

func (*Client) ContainerRename
func (cli *Client) ContainerRename(ctx context.Context, containerID, newContainerName string) error
ContainerRename changes the name of a given container.

func (*Client) ContainerResize
func (cli *Client) ContainerResize(ctx context.Context, containerID string, options types.ResizeOptions) error
ContainerResize changes the size of the tty for a container.

func (*Client) ContainerRestart
func (cli *Client) ContainerRestart(ctx context.Context, containerID string, timeout *time.Duration) error
ContainerRestart stops and starts a container again. It makes the daemon to wait for the container to be up again for a specific amount of time, given the timeout.

func (*Client) ContainerStart
func (cli *Client) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error
ContainerStart sends a request to the docker daemon to start a container.

func (*Client) ContainerStatPath
func (cli *Client) ContainerStatPath(ctx context.Context, containerID, path string) (types.ContainerPathStat, error)
ContainerStatPath returns Stat information about a path inside the container filesystem.

func (*Client) ContainerStats
func (cli *Client) ContainerStats(ctx context.Context, containerID string, stream bool) (types.ContainerStats, error)
ContainerStats returns near realtime stats for a given container. It's up to the caller to close the io.ReadCloser returned.

func (*Client) ContainerStop
func (cli *Client) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
ContainerStop stops a container without terminating the process. The process is blocked until the container stops or the timeout expires.

func (*Client) ContainerTop
func (cli *Client) ContainerTop(ctx context.Context, containerID string, arguments []string) (container.ContainerTopOKBody, error)
ContainerTop shows process information from within a container.

func (*Client) ContainerUnpause
func (cli *Client) ContainerUnpause(ctx context.Context, containerID string) error
ContainerUnpause resumes the process execution within a container

func (*Client) ContainerUpdate
func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error)
ContainerUpdate updates resources of a container

func (*Client) ContainerWait
func (cli *Client) ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error)
ContainerWait waits until the specified container is in a certain state indicated by the given condition, either "not-running" (default), "next-exit", or "removed".

If this client's API version is before 1.30, condition is ignored and ContainerWait will return immediately with the two channels, as the server will wait as if the condition were "not-running".

If this client's API version is at least 1.30, ContainerWait blocks until the request has been acknowledged by the server (with a response header), then returns two channels on which the caller can wait for the exit status of the container or an error if there was a problem either beginning the wait request or in getting the response. This allows the caller to synchronize ContainerWait with other calls, such as specifying a "next-exit" condition before issuing a ContainerStart request.

Example (WithTimeout)
func (*Client) ContainersPrune
func (cli *Client) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
ContainersPrune requests the daemon to delete unused data

func (*Client) CopyFromContainer
func (cli *Client) CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error)
CopyFromContainer gets the content from the container and returns it as a Reader to manipulate it in the host. It's up to the caller to close the reader.

func (*Client) CopyToContainer
func (cli *Client) CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error
CopyToContainer copies content into the container filesystem. Note that `content` must be a Reader for a TAR

func (*Client) CustomHTTPHeaders
func (cli *Client) CustomHTTPHeaders() map[string]string
CustomHTTPHeaders returns the custom http headers stored by the client.

func (*Client) DaemonHost
func (cli *Client) DaemonHost() string
DaemonHost returns the host address used by the client

func (*Client) DialSession
func (cli *Client) DialSession(ctx context.Context, proto string, meta map[string][]string) (net.Conn, error)
DialSession returns a connection that can be used communication with daemon

func (*Client) DiskUsage
func (cli *Client) DiskUsage(ctx context.Context) (types.DiskUsage, error)
DiskUsage requests the current data usage from the daemon

func (*Client) DistributionInspect
func (cli *Client) DistributionInspect(ctx context.Context, image, encodedRegistryAuth string) (registrytypes.DistributionInspect, error)
DistributionInspect returns the image digest with full Manifest

func (*Client) Events
func (cli *Client) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
Events returns a stream of events in the daemon. It's up to the caller to close the stream by cancelling the context. Once the stream has been completely read an io.EOF error will be sent over the error channel. If an error is sent all processing will be stopped. It's up to the caller to reopen the stream in the event of an error by reinvoking this method.

func (*Client) ImageBuild
func (cli *Client) ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
ImageBuild sends request to the daemon to build images. The Body in the response implement an io.ReadCloser and it's up to the caller to close it.

func (*Client) ImageCreate
func (cli *Client) ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error)
ImageCreate creates a new image based in the parent options. It returns the JSON content in the response body.

func (*Client) ImageHistory
func (cli *Client) ImageHistory(ctx context.Context, imageID string) ([]image.HistoryResponseItem, error)
ImageHistory returns the changes in an image in history format.

func (*Client) ImageImport
func (cli *Client) ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error)
ImageImport creates a new image based in the source options. It returns the JSON content in the response body.

func (*Client) ImageInspectWithRaw
func (cli *Client) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error)
ImageInspectWithRaw returns the image information and its raw representation.

func (*Client) ImageList
func (cli *Client) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
ImageList returns a list of images in the docker host.

func (*Client) ImageLoad
func (cli *Client) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)
ImageLoad loads an image in the docker host from the client host. It's up to the caller to close the io.ReadCloser in the ImageLoadResponse returned by this function.

func (*Client) ImagePull
func (cli *Client) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)
ImagePull requests the docker host to pull an image from a remote registry. It executes the privileged function if the operation is unauthorized and it tries one more time. It's up to the caller to handle the io.ReadCloser and close it properly.

FIXME(vdemeester): there is currently used in a few way in docker/docker - if not in trusted content, ref is used to pass the whole reference, and tag is empty - if in trusted content, ref is used to pass the reference name, and tag for the digest

func (*Client) ImagePush
func (cli *Client) ImagePush(ctx context.Context, image string, options types.ImagePushOptions) (io.ReadCloser, error)
ImagePush requests the docker host to push an image to a remote registry. It executes the privileged function if the operation is unauthorized and it tries one more time. It's up to the caller to handle the io.ReadCloser and close it properly.

func (*Client) ImageRemove
func (cli *Client) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)
ImageRemove removes an image from the docker host.

func (*Client) ImageSave
func (cli *Client) ImageSave(ctx context.Context, imageIDs []string) (io.ReadCloser, error)
ImageSave retrieves one or more images from the docker host as an io.ReadCloser. It's up to the caller to store the images and close the stream.

func (*Client) ImageSearch
func (cli *Client) ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error)
ImageSearch makes the docker host to search by a term in a remote registry. The list of results is not sorted in any fashion.

func (*Client) ImageTag
func (cli *Client) ImageTag(ctx context.Context, source, target string) error
ImageTag tags an image in the docker host

func (*Client) ImagesPrune
func (cli *Client) ImagesPrune(ctx context.Context, pruneFilters filters.Args) (types.ImagesPruneReport, error)
ImagesPrune requests the daemon to delete unused data

func (*Client) Info
func (cli *Client) Info(ctx context.Context) (types.Info, error)
Info returns information about the docker server.

func (*Client) NegotiateAPIVersion
func (cli *Client) NegotiateAPIVersion(ctx context.Context)
NegotiateAPIVersion queries the API and updates the version to match the API version. Any errors are silently ignored.

func (*Client) NegotiateAPIVersionPing
func (cli *Client) NegotiateAPIVersionPing(p types.Ping)
NegotiateAPIVersionPing updates the client version to match the Ping.APIVersion if the ping version is less than the default version.

func (*Client) NetworkConnect
func (cli *Client) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error
NetworkConnect connects a container to an existent network in the docker host.

func (*Client) NetworkCreate
func (cli *Client) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)
NetworkCreate creates a new network in the docker host.

func (*Client) NetworkDisconnect
func (cli *Client) NetworkDisconnect(ctx context.Context, networkID, containerID string, force bool) error
NetworkDisconnect disconnects a container from an existent network in the docker host.

func (*Client) NetworkInspect
func (cli *Client) NetworkInspect(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, error)
NetworkInspect returns the information for a specific network configured in the docker host.

func (*Client) NetworkInspectWithRaw
func (cli *Client) NetworkInspectWithRaw(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error)
NetworkInspectWithRaw returns the information for a specific network configured in the docker host and its raw representation.

func (*Client) NetworkList
func (cli *Client) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error)
NetworkList returns the list of networks configured in the docker host.

func (*Client) NetworkRemove
func (cli *Client) NetworkRemove(ctx context.Context, networkID string) error
NetworkRemove removes an existent network from the docker host.

func (*Client) NetworksPrune
func (cli *Client) NetworksPrune(ctx context.Context, pruneFilters filters.Args) (types.NetworksPruneReport, error)
NetworksPrune requests the daemon to delete unused networks

func (*Client) NewVersionError
func (cli *Client) NewVersionError(APIrequired, feature string) error
NewVersionError returns an error if the APIVersion required if less than the current supported version

func (*Client) NodeInspectWithRaw
func (cli *Client) NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error)
NodeInspectWithRaw returns the node information.

func (*Client) NodeList
func (cli *Client) NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error)
NodeList returns the list of nodes.

func (*Client) NodeRemove
func (cli *Client) NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error
NodeRemove removes a Node.

func (*Client) NodeUpdate
func (cli *Client) NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error
NodeUpdate updates a Node.

func (*Client) Ping
func (cli *Client) Ping(ctx context.Context) (types.Ping, error)
Ping pings the server and returns the value of the "Docker-Experimental", "OS-Type" & "API-Version" headers

func (*Client) PluginCreate
func (cli *Client) PluginCreate(ctx context.Context, createContext io.Reader, createOptions types.PluginCreateOptions) error
PluginCreate creates a plugin

func (*Client) PluginDisable
func (cli *Client) PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error
PluginDisable disables a plugin

func (*Client) PluginEnable
func (cli *Client) PluginEnable(ctx context.Context, name string, options types.PluginEnableOptions) error
PluginEnable enables a plugin

func (*Client) PluginInspectWithRaw
func (cli *Client) PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error)
PluginInspectWithRaw inspects an existing plugin

func (*Client) PluginInstall
func (cli *Client) PluginInstall(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
PluginInstall installs a plugin

func (*Client) PluginList
func (cli *Client) PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error)
PluginList returns the installed plugins

func (*Client) PluginPush
func (cli *Client) PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error)
PluginPush pushes a plugin to a registry

func (*Client) PluginRemove
func (cli *Client) PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error
PluginRemove removes a plugin

func (*Client) PluginSet
func (cli *Client) PluginSet(ctx context.Context, name string, args []string) error
PluginSet modifies settings for an existing plugin

func (*Client) PluginUpgrade
func (cli *Client) PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
PluginUpgrade upgrades a plugin

func (*Client) RegistryLogin
func (cli *Client) RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)
RegistryLogin authenticates the docker server with a given docker registry. It returns unauthorizedError when the authentication fails.

func (*Client) SecretCreate
func (cli *Client) SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error)
SecretCreate creates a new Secret.

func (*Client) SecretInspectWithRaw
func (cli *Client) SecretInspectWithRaw(ctx context.Context, id string) (swarm.Secret, []byte, error)
SecretInspectWithRaw returns the secret information with raw data

func (*Client) SecretList
func (cli *Client) SecretList(ctx context.Context, options types.SecretListOptions) ([]swarm.Secret, error)
SecretList returns the list of secrets.

func (*Client) SecretRemove
func (cli *Client) SecretRemove(ctx context.Context, id string) error
SecretRemove removes a Secret.

func (*Client) SecretUpdate
func (cli *Client) SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error
SecretUpdate attempts to update a Secret

func (*Client) ServerVersion
func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error)
ServerVersion returns information of the docker client and server host.

func (*Client) ServiceCreate
func (cli *Client) ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error)
ServiceCreate creates a new Service.

func (*Client) ServiceInspectWithRaw
func (cli *Client) ServiceInspectWithRaw(ctx context.Context, serviceID string, opts types.ServiceInspectOptions) (swarm.Service, []byte, error)
ServiceInspectWithRaw returns the service information and the raw data.

func (*Client) ServiceList
func (cli *Client) ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
ServiceList returns the list of services.

func (*Client) ServiceLogs
func (cli *Client) ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
ServiceLogs returns the logs generated by a service in an io.ReadCloser. It's up to the caller to close the stream.

Example (WithTimeout)
func (*Client) ServiceRemove
func (cli *Client) ServiceRemove(ctx context.Context, serviceID string) error
ServiceRemove kills and removes a service.

func (*Client) ServiceUpdate
func (cli *Client) ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) (types.ServiceUpdateResponse, error)
ServiceUpdate updates a Service.

func (*Client) SetCustomHTTPHeaders
func (cli *Client) SetCustomHTTPHeaders(headers map[string]string)
SetCustomHTTPHeaders that will be set on every HTTP request made by the client.

func (*Client) SwarmGetUnlockKey
func (cli *Client) SwarmGetUnlockKey(ctx context.Context) (types.SwarmUnlockKeyResponse, error)
SwarmGetUnlockKey retrieves the swarm's unlock key.

func (*Client) SwarmInit
func (cli *Client) SwarmInit(ctx context.Context, req swarm.InitRequest) (string, error)
SwarmInit initializes the swarm.

func (*Client) SwarmInspect
func (cli *Client) SwarmInspect(ctx context.Context) (swarm.Swarm, error)
SwarmInspect inspects the swarm.

func (*Client) SwarmJoin
func (cli *Client) SwarmJoin(ctx context.Context, req swarm.JoinRequest) error
SwarmJoin joins the swarm.

func (*Client) SwarmLeave
func (cli *Client) SwarmLeave(ctx context.Context, force bool) error
SwarmLeave leaves the swarm.

func (*Client) SwarmUnlock
func (cli *Client) SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error
SwarmUnlock unlocks locked swarm.

func (*Client) SwarmUpdate
func (cli *Client) SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error
SwarmUpdate updates the swarm.

func (*Client) TaskInspectWithRaw
func (cli *Client) TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error)
TaskInspectWithRaw returns the task information and its raw representation..

func (*Client) TaskList
func (cli *Client) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error)
TaskList returns the list of tasks.

func (*Client) TaskLogs
func (cli *Client) TaskLogs(ctx context.Context, taskID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
TaskLogs returns the logs generated by a task in an io.ReadCloser. It's up to the caller to close the stream.

func (*Client) VolumeCreate
func (cli *Client) VolumeCreate(ctx context.Context, options volumetypes.VolumesCreateBody) (types.Volume, error)
VolumeCreate creates a volume in the docker host.

func (*Client) VolumeInspect
func (cli *Client) VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error)
VolumeInspect returns the information about a specific volume in the docker host.

func (*Client) VolumeInspectWithRaw
func (cli *Client) VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error)
VolumeInspectWithRaw returns the information about a specific volume in the docker host and its raw representation

func (*Client) VolumeList
func (cli *Client) VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumesListOKBody, error)
VolumeList returns the volumes configured in the docker host.

func (*Client) VolumeRemove
func (cli *Client) VolumeRemove(ctx context.Context, volumeID string, force bool) error
VolumeRemove removes a volume from the docker host.

func (*Client) VolumesPrune
func (cli *Client) VolumesPrune(ctx context.Context, pruneFilters filters.Args) (types.VolumesPruneReport, error)
VolumesPrune requests the daemon to delete unused data

type CommonAPIClient
type CommonAPIClient interface {
    ConfigAPIClient
    ContainerAPIClient
    DistributionAPIClient
    ImageAPIClient
    NodeAPIClient
    NetworkAPIClient
    PluginAPIClient
    ServiceAPIClient
    SwarmAPIClient
    SecretAPIClient
    SystemAPIClient
    VolumeAPIClient
    ClientVersion() string
    DaemonHost() string
    ServerVersion(ctx context.Context) (types.Version, error)
    NegotiateAPIVersion(ctx context.Context)
    NegotiateAPIVersionPing(types.Ping)
    DialSession(ctx context.Context, proto string, meta map[string][]string) (net.Conn, error)
}
CommonAPIClient is the common methods between stable and experimental versions of APIClient.

type ConfigAPIClient
type ConfigAPIClient interface {
    ConfigList(ctx context.Context, options types.ConfigListOptions) ([]swarm.Config, error)
    ConfigCreate(ctx context.Context, config swarm.ConfigSpec) (types.ConfigCreateResponse, error)
    ConfigRemove(ctx context.Context, id string) error
    ConfigInspectWithRaw(ctx context.Context, name string) (swarm.Config, []byte, error)
    ConfigUpdate(ctx context.Context, id string, version swarm.Version, config swarm.ConfigSpec) error
}
ConfigAPIClient defines API client methods for configs

type ContainerAPIClient
type ContainerAPIClient interface {
    ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
    ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
    ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
    ContainerDiff(ctx context.Context, container string) ([]container.ContainerChangeResponseItem, error)
    ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error)
    ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
    ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
    ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
    ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
    ContainerExport(ctx context.Context, container string) (io.ReadCloser, error)
    ContainerInspect(ctx context.Context, container string) (types.ContainerJSON, error)
    ContainerInspectWithRaw(ctx context.Context, container string, getSize bool) (types.ContainerJSON, []byte, error)
    ContainerKill(ctx context.Context, container, signal string) error
    ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
    ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
    ContainerPause(ctx context.Context, container string) error
    ContainerRemove(ctx context.Context, container string, options types.ContainerRemoveOptions) error
    ContainerRename(ctx context.Context, container, newContainerName string) error
    ContainerResize(ctx context.Context, container string, options types.ResizeOptions) error
    ContainerRestart(ctx context.Context, container string, timeout *time.Duration) error
    ContainerStatPath(ctx context.Context, container, path string) (types.ContainerPathStat, error)
    ContainerStats(ctx context.Context, container string, stream bool) (types.ContainerStats, error)
    ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error
    ContainerStop(ctx context.Context, container string, timeout *time.Duration) error
    ContainerTop(ctx context.Context, container string, arguments []string) (container.ContainerTopOKBody, error)
    ContainerUnpause(ctx context.Context, container string) error
    ContainerUpdate(ctx context.Context, container string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error)
    ContainerWait(ctx context.Context, container string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error)
    CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error)
    CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error
    ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
}
ContainerAPIClient defines API client methods for the containers

type DistributionAPIClient
type DistributionAPIClient interface {
    DistributionInspect(ctx context.Context, image, encodedRegistryAuth string) (registry.DistributionInspect, error)
}
DistributionAPIClient defines API client methods for the registry

type ImageAPIClient
type ImageAPIClient interface {
    ImageBuild(ctx context.Context, context io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
    BuildCachePrune(ctx context.Context) (*types.BuildCachePruneReport, error)
    ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error)
    ImageHistory(ctx context.Context, image string) ([]image.HistoryResponseItem, error)
    ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error)
    ImageInspectWithRaw(ctx context.Context, image string) (types.ImageInspect, []byte, error)
    ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
    ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)
    ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error)
    ImagePush(ctx context.Context, ref string, options types.ImagePushOptions) (io.ReadCloser, error)
    ImageRemove(ctx context.Context, image string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)
    ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error)
    ImageSave(ctx context.Context, images []string) (io.ReadCloser, error)
    ImageTag(ctx context.Context, image, ref string) error
    ImagesPrune(ctx context.Context, pruneFilter filters.Args) (types.ImagesPruneReport, error)
}
ImageAPIClient defines API client methods for the images

type NetworkAPIClient
type NetworkAPIClient interface {
    NetworkConnect(ctx context.Context, networkID, container string, config *network.EndpointSettings) error
    NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)
    NetworkDisconnect(ctx context.Context, networkID, container string, force bool) error
    NetworkInspect(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, error)
    NetworkInspectWithRaw(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error)
    NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error)
    NetworkRemove(ctx context.Context, networkID string) error
    NetworksPrune(ctx context.Context, pruneFilter filters.Args) (types.NetworksPruneReport, error)
}
NetworkAPIClient defines API client methods for the networks

type NodeAPIClient
type NodeAPIClient interface {
    NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error)
    NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error)
    NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error
    NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error
}
NodeAPIClient defines API client methods for the nodes

type PluginAPIClient
type PluginAPIClient interface {
    PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error)
    PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error
    PluginEnable(ctx context.Context, name string, options types.PluginEnableOptions) error
    PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error
    PluginInstall(ctx context.Context, name string, options types.PluginInstallOptions) (io.ReadCloser, error)
    PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (io.ReadCloser, error)
    PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error)
    PluginSet(ctx context.Context, name string, args []string) error
    PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error)
    PluginCreate(ctx context.Context, createContext io.Reader, options types.PluginCreateOptions) error
}
PluginAPIClient defines API client methods for the plugins

type SecretAPIClient
type SecretAPIClient interface {
    SecretList(ctx context.Context, options types.SecretListOptions) ([]swarm.Secret, error)
    SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error)
    SecretRemove(ctx context.Context, id string) error
    SecretInspectWithRaw(ctx context.Context, name string) (swarm.Secret, []byte, error)
    SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error
}
SecretAPIClient defines API client methods for secrets

type ServiceAPIClient
type ServiceAPIClient interface {
    ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error)
    ServiceInspectWithRaw(ctx context.Context, serviceID string, options types.ServiceInspectOptions) (swarm.Service, []byte, error)
    ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
    ServiceRemove(ctx context.Context, serviceID string) error
    ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) (types.ServiceUpdateResponse, error)
    ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
    TaskLogs(ctx context.Context, taskID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
    TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error)
    TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error)
}
ServiceAPIClient defines API client methods for the services

type SwarmAPIClient
type SwarmAPIClient interface {
    SwarmInit(ctx context.Context, req swarm.InitRequest) (string, error)
    SwarmJoin(ctx context.Context, req swarm.JoinRequest) error
    SwarmGetUnlockKey(ctx context.Context) (types.SwarmUnlockKeyResponse, error)
    SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error
    SwarmLeave(ctx context.Context, force bool) error
    SwarmInspect(ctx context.Context) (swarm.Swarm, error)
    SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error
}
SwarmAPIClient defines API client methods for the swarm

type SystemAPIClient
type SystemAPIClient interface {
    Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
    Info(ctx context.Context) (types.Info, error)
    RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)
    DiskUsage(ctx context.Context) (types.DiskUsage, error)
    Ping(ctx context.Context) (types.Ping, error)
}
SystemAPIClient defines API client methods for the system

type VolumeAPIClient
type VolumeAPIClient interface {
    VolumeCreate(ctx context.Context, options volumetypes.VolumesCreateBody) (types.Volume, error)
    VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error)
    VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error)
    VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumesListOKBody, error)
    VolumeRemove(ctx context.Context, volumeID string, force bool) error
    VolumesPrune(ctx context.Context, pruneFilter filters.Args) (types.VolumesPruneReport, error)
}
VolumeAPIClient defines API client methods for the volumes

Directories
Path	Synopsis
api	
api/types	Package types is used for API stability in the types and response to the consumers of the API stats endpoint.
api/types/blkiodev	
api/types/container	
api/types/events	
api/types/filters	Package filters provides tools for encoding a mapping of keys to a set of multiple values.
api/types/image	
api/types/mount	
api/types/network	
api/types/registry	
api/types/strslice	
api/types/swarm	
api/types/swarm/runtime	Package runtime is a generated protocol buffer package.
api/types/time	
api/types/versions	
api/types/versions/v1p19	Package v1p19 provides specific API types for the API version 1, patch 19.
api/types/versions/v1p20	Package v1p20 provides specific API types for the API version 1, patch 20.
api/types/volume	
Package docker imports 39 packages (graph) and is imported by 21 packages. Updated 9 months ago. Refresh now. Tools for package owners.

Website Issues | Go Language Back to top
