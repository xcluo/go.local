2:24:func main() {
3:30:	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
4:42:func CheckRedirect(req *http.Request, via []*http.Request) error
5:43:func ErrorConnectionFailed(host string) error
6:44:func IsErrConfigNotFound(err error) bool
7:45:func IsErrConnectionFailed(err error) bool
8:46:func IsErrContainerNotFound(err error) bool
9:47:func IsErrImageNotFound(err error) bool
10:48:func IsErrNetworkNotFound(err error) bool
11:49:func IsErrNodeNotFound(err error) bool
12:50:func IsErrNotFound(err error) bool
13:51:func IsErrNotImplemented(err error) bool
14:52:func IsErrPluginNotFound(err error) bool
15:53:func IsErrPluginPermissionDenied(err error) bool
16:54:func IsErrSecretNotFound(err error) bool
17:55:func IsErrServiceNotFound(err error) bool
18:56:func IsErrTaskNotFound(err error) bool
19:57:func IsErrUnauthorized(err error) bool
20:58:func IsErrVolumeNotFound(err error) bool
21:59:func ParseHost(host string) (string, string, string, error)
22:60:func ParseHostURL(host string) (*url.URL, error)
23:61:type APIClient
24:62:type CheckpointAPIClient
25:63:type Client
26:64:func NewClient(host string, version string, client *http.Client, httpHeaders map[string]string) (*Client, error)
27:65:func NewEnvClient() (*Client, error)
28:66:func (cli *Client) BuildCachePrune(ctx context.Context) (*types.BuildCachePruneReport, error)
29:67:func (cli *Client) CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) error
30:68:func (cli *Client) CheckpointDelete(ctx context.Context, containerID string, options types.CheckpointDeleteOptions) error
31:69:func (cli *Client) CheckpointList(ctx context.Context, container string, options types.CheckpointListOptions) ([]types.Checkpoint, error)
32:70:func (cli *Client) ClientVersion() string
33:71:func (cli *Client) Close() error
34:72:func (cli *Client) ConfigCreate(ctx context.Context, config swarm.ConfigSpec) (types.ConfigCreateResponse, error)
35:73:func (cli *Client) ConfigInspectWithRaw(ctx context.Context, id string) (swarm.Config, []byte, error)
36:74:func (cli *Client) ConfigList(ctx context.Context, options types.ConfigListOptions) ([]swarm.Config, error)
37:75:func (cli *Client) ConfigRemove(ctx context.Context, id string) error
38:76:func (cli *Client) ConfigUpdate(ctx context.Context, id string, version swarm.Version, config swarm.ConfigSpec) error
39:77:func (cli *Client) ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
40:78:func (cli *Client) ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
41:79:func (cli *Client) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
42:80:func (cli *Client) ContainerDiff(ctx context.Context, containerID string) ([]container.ContainerChangeResponseItem, error)
43:81:func (cli *Client) ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error)
44:82:func (cli *Client) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
45:83:func (cli *Client) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
46:84:func (cli *Client) ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
47:85:func (cli *Client) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
48:86:func (cli *Client) ContainerExport(ctx context.Context, containerID string) (io.ReadCloser, error)
49:87:func (cli *Client) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
50:88:func (cli *Client) ContainerInspectWithRaw(ctx context.Context, containerID string, getSize bool) (types.ContainerJSON, []byte, error)
51:89:func (cli *Client) ContainerKill(ctx context.Context, containerID, signal string) error
52:90:func (cli *Client) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
53:91:func (cli *Client) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
54:92:func (cli *Client) ContainerPause(ctx context.Context, containerID string) error
55:93:func (cli *Client) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
56:94:func (cli *Client) ContainerRename(ctx context.Context, containerID, newContainerName string) error
57:95:func (cli *Client) ContainerResize(ctx context.Context, containerID string, options types.ResizeOptions) error
58:96:func (cli *Client) ContainerRestart(ctx context.Context, containerID string, timeout *time.Duration) error
59:97:func (cli *Client) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error
60:98:func (cli *Client) ContainerStatPath(ctx context.Context, containerID, path string) (types.ContainerPathStat, error)
61:99:func (cli *Client) ContainerStats(ctx context.Context, containerID string, stream bool) (types.ContainerStats, error)
62:100:func (cli *Client) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
63:101:func (cli *Client) ContainerTop(ctx context.Context, containerID string, arguments []string) (container.ContainerTopOKBody, error)
64:102:func (cli *Client) ContainerUnpause(ctx context.Context, containerID string) error
65:103:func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error)
66:104:func (cli *Client) ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error)
67:105:func (cli *Client) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
68:106:func (cli *Client) CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error)
69:107:func (cli *Client) CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error
70:108:func (cli *Client) CustomHTTPHeaders() map[string]string
71:109:func (cli *Client) DaemonHost() string
72:110:func (cli *Client) DialSession(ctx context.Context, proto string, meta map[string][]string) (net.Conn, error)
73:111:func (cli *Client) DiskUsage(ctx context.Context) (types.DiskUsage, error)
74:112:func (cli *Client) DistributionInspect(ctx context.Context, image, encodedRegistryAuth string) (registrytypes.DistributionInspect, error)
75:113:func (cli *Client) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
76:114:func (cli *Client) ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
77:115:func (cli *Client) ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error)
78:116:func (cli *Client) ImageHistory(ctx context.Context, imageID string) ([]image.HistoryResponseItem, error)
79:117:func (cli *Client) ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error)
80:118:func (cli *Client) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error)
81:119:func (cli *Client) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
82:120:func (cli *Client) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)
83:121:func (cli *Client) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)
84:122:func (cli *Client) ImagePush(ctx context.Context, image string, options types.ImagePushOptions) (io.ReadCloser, error)
85:123:func (cli *Client) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)
86:124:func (cli *Client) ImageSave(ctx context.Context, imageIDs []string) (io.ReadCloser, error)
87:125:func (cli *Client) ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error)
88:126:func (cli *Client) ImageTag(ctx context.Context, source, target string) error
89:127:func (cli *Client) ImagesPrune(ctx context.Context, pruneFilters filters.Args) (types.ImagesPruneReport, error)
90:128:func (cli *Client) Info(ctx context.Context) (types.Info, error)
91:129:func (cli *Client) NegotiateAPIVersion(ctx context.Context)
92:130:func (cli *Client) NegotiateAPIVersionPing(p types.Ping)
93:131:func (cli *Client) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error
94:132:func (cli *Client) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)
95:133:func (cli *Client) NetworkDisconnect(ctx context.Context, networkID, containerID string, force bool) error
96:134:func (cli *Client) NetworkInspect(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, error)
97:135:func (cli *Client) NetworkInspectWithRaw(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error)
98:136:func (cli *Client) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error)
99:137:func (cli *Client) NetworkRemove(ctx context.Context, networkID string) error
100:138:func (cli *Client) NetworksPrune(ctx context.Context, pruneFilters filters.Args) (types.NetworksPruneReport, error)
101:139:func (cli *Client) NewVersionError(APIrequired, feature string) error
102:140:func (cli *Client) NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error)
103:141:func (cli *Client) NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error)
104:142:func (cli *Client) NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error
105:143:func (cli *Client) NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error
106:144:func (cli *Client) Ping(ctx context.Context) (types.Ping, error)
107:145:func (cli *Client) PluginCreate(ctx context.Context, createContext io.Reader, createOptions types.PluginCreateOptions) error
108:146:func (cli *Client) PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error
109:147:func (cli *Client) PluginEnable(ctx context.Context, name string, options types.PluginEnableOptions) error
110:148:func (cli *Client) PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error)
111:149:func (cli *Client) PluginInstall(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
112:150:func (cli *Client) PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error)
113:151:func (cli *Client) PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error)
114:152:func (cli *Client) PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error
115:153:func (cli *Client) PluginSet(ctx context.Context, name string, args []string) error
116:154:func (cli *Client) PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
117:155:func (cli *Client) RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)
118:156:func (cli *Client) SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error)
119:157:func (cli *Client) SecretInspectWithRaw(ctx context.Context, id string) (swarm.Secret, []byte, error)
120:158:func (cli *Client) SecretList(ctx context.Context, options types.SecretListOptions) ([]swarm.Secret, error)
121:159:func (cli *Client) SecretRemove(ctx context.Context, id string) error
122:160:func (cli *Client) SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error
123:161:func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error)
124:162:func (cli *Client) ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error)
125:163:func (cli *Client) ServiceInspectWithRaw(ctx context.Context, serviceID string, opts types.ServiceInspectOptions) (swarm.Service, []byte, error)
126:164:func (cli *Client) ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
127:165:func (cli *Client) ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
128:166:func (cli *Client) ServiceRemove(ctx context.Context, serviceID string) error
129:167:func (cli *Client) ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) (types.ServiceUpdateResponse, error)
130:168:func (cli *Client) SetCustomHTTPHeaders(headers map[string]string)
131:169:func (cli *Client) SwarmGetUnlockKey(ctx context.Context) (types.SwarmUnlockKeyResponse, error)
132:170:func (cli *Client) SwarmInit(ctx context.Context, req swarm.InitRequest) (string, error)
133:171:func (cli *Client) SwarmInspect(ctx context.Context) (swarm.Swarm, error)
134:172:func (cli *Client) SwarmJoin(ctx context.Context, req swarm.JoinRequest) error
135:173:func (cli *Client) SwarmLeave(ctx context.Context, force bool) error
136:174:func (cli *Client) SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error
137:175:func (cli *Client) SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error
138:176:func (cli *Client) TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error)
139:177:func (cli *Client) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error)
140:178:func (cli *Client) TaskLogs(ctx context.Context, taskID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
141:179:func (cli *Client) VolumeCreate(ctx context.Context, options volumetypes.VolumesCreateBody) (types.Volume, error)
142:180:func (cli *Client) VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error)
143:181:func (cli *Client) VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error)
144:182:func (cli *Client) VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumesListOKBody, error)
145:183:func (cli *Client) VolumeRemove(ctx context.Context, volumeID string, force bool) error
146:184:func (cli *Client) VolumesPrune(ctx context.Context, pruneFilters filters.Args) (types.VolumesPruneReport, error)
147:185:type CommonAPIClient
148:186:type ConfigAPIClient
149:187:type ContainerAPIClient
150:188:type DistributionAPIClient
151:189:type ImageAPIClient
152:190:type NetworkAPIClient
153:191:type NodeAPIClient
154:192:type PluginAPIClient
155:193:type SecretAPIClient
156:194:type ServiceAPIClient
157:195:type SwarmAPIClient
158:196:type SystemAPIClient
159:197:type VolumeAPIClient
160:213:func CheckRedirect
161:214:func CheckRedirect(req *http.Request, via []*http.Request) error
162:219:func ErrorConnectionFailed
163:220:func ErrorConnectionFailed(host string) error
164:223:func IsErrConfigNotFound
165:224:func IsErrConfigNotFound(err error) bool
166:229:func IsErrConnectionFailed
167:230:func IsErrConnectionFailed(err error) bool
168:233:func IsErrContainerNotFound
169:234:func IsErrContainerNotFound(err error) bool
170:239:func IsErrImageNotFound
171:240:func IsErrImageNotFound(err error) bool
172:245:func IsErrNetworkNotFound
173:246:func IsErrNetworkNotFound(err error) bool
174:251:func IsErrNodeNotFound
175:252:func IsErrNodeNotFound(err error) bool
176:257:func IsErrNotFound
177:258:func IsErrNotFound(err error) bool
178:261:func IsErrNotImplemented
179:262:func IsErrNotImplemented(err error) bool
180:265:func IsErrPluginNotFound
181:266:func IsErrPluginNotFound(err error) bool
182:271:func IsErrPluginPermissionDenied
183:272:func IsErrPluginPermissionDenied(err error) bool
184:275:func IsErrSecretNotFound
185:276:func IsErrSecretNotFound(err error) bool
186:281:func IsErrServiceNotFound
187:282:func IsErrServiceNotFound(err error) bool
188:287:func IsErrTaskNotFound
189:288:func IsErrTaskNotFound(err error) bool
190:293:func IsErrUnauthorized
191:294:func IsErrUnauthorized(err error) bool
192:297:func IsErrVolumeNotFound
193:298:func IsErrVolumeNotFound(err error) bool
194:303:func ParseHost
195:304:func ParseHost(host string) (string, string, string, error)
196:307:func ParseHostURL
197:308:func ParseHostURL(host string) (*url.URL, error)
198:311:type APIClient
199:312:type APIClient interface {
200:318:type CheckpointAPIClient
201:319:type CheckpointAPIClient interface {
202:320:    CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) error
203:321:    CheckpointDelete(ctx context.Context, container string, options types.CheckpointDeleteOptions) error
204:322:    CheckpointList(ctx context.Context, container string, options types.CheckpointListOptions) ([]types.Checkpoint, error)
205:326:type Client
206:327:type Client struct {
207:332:func NewClient
208:333:func NewClient(host string, version string, client *http.Client, httpHeaders map[string]string) (*Client, error)
209:338:func NewEnvClient
210:339:func NewEnvClient() (*Client, error)
211:342:func (*Client) BuildCachePrune
212:343:func (cli *Client) BuildCachePrune(ctx context.Context) (*types.BuildCachePruneReport, error)
213:346:func (*Client) CheckpointCreate
214:347:func (cli *Client) CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) error
215:350:func (*Client) CheckpointDelete
216:351:func (cli *Client) CheckpointDelete(ctx context.Context, containerID string, options types.CheckpointDeleteOptions) error
217:354:func (*Client) CheckpointList
218:355:func (cli *Client) CheckpointList(ctx context.Context, container string, options types.CheckpointListOptions) ([]types.Checkpoint, error)
219:358:func (*Client) ClientVersion
220:359:func (cli *Client) ClientVersion() string
221:362:func (*Client) Close
222:363:func (cli *Client) Close() error
223:366:func (*Client) ConfigCreate
224:367:func (cli *Client) ConfigCreate(ctx context.Context, config swarm.ConfigSpec) (types.ConfigCreateResponse, error)
225:370:func (*Client) ConfigInspectWithRaw
226:371:func (cli *Client) ConfigInspectWithRaw(ctx context.Context, id string) (swarm.Config, []byte, error)
227:374:func (*Client) ConfigList
228:375:func (cli *Client) ConfigList(ctx context.Context, options types.ConfigListOptions) ([]swarm.Config, error)
229:378:func (*Client) ConfigRemove
230:379:func (cli *Client) ConfigRemove(ctx context.Context, id string) error
231:382:func (*Client) ConfigUpdate
232:383:func (cli *Client) ConfigUpdate(ctx context.Context, id string, version swarm.Version, config swarm.ConfigSpec) error
233:386:func (*Client) ContainerAttach
234:387:func (cli *Client) ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
235:388:ContainerAttach attaches a connection to a container in the server. It returns a types.HijackedConnection with the hijacked connection and the a reader to get output. It's up to the called to close the hijacked connection by calling types.HijackedResponse.Close.
236:403:func (*Client) ContainerCommit
237:404:func (cli *Client) ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
238:407:func (*Client) ContainerCreate
239:408:func (cli *Client) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
240:411:func (*Client) ContainerDiff
241:412:func (cli *Client) ContainerDiff(ctx context.Context, containerID string) ([]container.ContainerChangeResponseItem, error)
242:415:func (*Client) ContainerExecAttach
243:416:func (cli *Client) ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error)
244:417:ContainerExecAttach attaches a connection to an exec process in the server. It returns a types.HijackedConnection with the hijacked connection and the a reader to get output. It's up to the called to close the hijacked connection by calling types.HijackedResponse.Close.
245:419:func (*Client) ContainerExecCreate
246:420:func (cli *Client) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
247:423:func (*Client) ContainerExecInspect
248:424:func (cli *Client) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
249:427:func (*Client) ContainerExecResize
250:428:func (cli *Client) ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
251:431:func (*Client) ContainerExecStart
252:432:func (cli *Client) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
253:435:func (*Client) ContainerExport
254:436:func (cli *Client) ContainerExport(ctx context.Context, containerID string) (io.ReadCloser, error)
255:439:func (*Client) ContainerInspect
256:440:func (cli *Client) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
257:443:func (*Client) ContainerInspectWithRaw
258:444:func (cli *Client) ContainerInspectWithRaw(ctx context.Context, containerID string, getSize bool) (types.ContainerJSON, []byte, error)
259:447:func (*Client) ContainerKill
260:448:func (cli *Client) ContainerKill(ctx context.Context, containerID, signal string) error
261:451:func (*Client) ContainerList
262:452:func (cli *Client) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
263:455:func (*Client) ContainerLogs
264:456:func (cli *Client) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
265:473:func (*Client) ContainerPause
266:474:func (cli *Client) ContainerPause(ctx context.Context, containerID string) error
267:477:func (*Client) ContainerRemove
268:478:func (cli *Client) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
269:481:func (*Client) ContainerRename
270:482:func (cli *Client) ContainerRename(ctx context.Context, containerID, newContainerName string) error
271:485:func (*Client) ContainerResize
272:486:func (cli *Client) ContainerResize(ctx context.Context, containerID string, options types.ResizeOptions) error
273:489:func (*Client) ContainerRestart
274:490:func (cli *Client) ContainerRestart(ctx context.Context, containerID string, timeout *time.Duration) error
275:493:func (*Client) ContainerStart
276:494:func (cli *Client) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error
277:497:func (*Client) ContainerStatPath
278:498:func (cli *Client) ContainerStatPath(ctx context.Context, containerID, path string) (types.ContainerPathStat, error)
279:501:func (*Client) ContainerStats
280:502:func (cli *Client) ContainerStats(ctx context.Context, containerID string, stream bool) (types.ContainerStats, error)
281:505:func (*Client) ContainerStop
282:506:func (cli *Client) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
283:509:func (*Client) ContainerTop
284:510:func (cli *Client) ContainerTop(ctx context.Context, containerID string, arguments []string) (container.ContainerTopOKBody, error)
285:513:func (*Client) ContainerUnpause
286:514:func (cli *Client) ContainerUnpause(ctx context.Context, containerID string) error
287:517:func (*Client) ContainerUpdate
288:518:func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error)
289:521:func (*Client) ContainerWait
290:522:func (cli *Client) ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error)
291:530:func (*Client) ContainersPrune
292:531:func (cli *Client) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
293:534:func (*Client) CopyFromContainer
294:535:func (cli *Client) CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error)
295:538:func (*Client) CopyToContainer
296:539:func (cli *Client) CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error
297:542:func (*Client) CustomHTTPHeaders
298:543:func (cli *Client) CustomHTTPHeaders() map[string]string
299:546:func (*Client) DaemonHost
300:547:func (cli *Client) DaemonHost() string
301:550:func (*Client) DialSession
302:551:func (cli *Client) DialSession(ctx context.Context, proto string, meta map[string][]string) (net.Conn, error)
303:554:func (*Client) DiskUsage
304:555:func (cli *Client) DiskUsage(ctx context.Context) (types.DiskUsage, error)
305:558:func (*Client) DistributionInspect
306:559:func (cli *Client) DistributionInspect(ctx context.Context, image, encodedRegistryAuth string) (registrytypes.DistributionInspect, error)
307:562:func (*Client) Events
308:563:func (cli *Client) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
309:566:func (*Client) ImageBuild
310:567:func (cli *Client) ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
311:570:func (*Client) ImageCreate
312:571:func (cli *Client) ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error)
313:574:func (*Client) ImageHistory
314:575:func (cli *Client) ImageHistory(ctx context.Context, imageID string) ([]image.HistoryResponseItem, error)
315:578:func (*Client) ImageImport
316:579:func (cli *Client) ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error)
317:582:func (*Client) ImageInspectWithRaw
318:583:func (cli *Client) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error)
319:586:func (*Client) ImageList
320:587:func (cli *Client) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
321:590:func (*Client) ImageLoad
322:591:func (cli *Client) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)
323:592:ImageLoad loads an image in the docker host from the client host. It's up to the caller to close the io.ReadCloser in the ImageLoadResponse returned by this function.
324:594:func (*Client) ImagePull
325:595:func (cli *Client) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)
326:596:ImagePull requests the docker host to pull an image from a remote registry. It executes the privileged function if the operation is unauthorized and it tries one more time. It's up to the caller to handle the io.ReadCloser and close it properly.
327:600:func (*Client) ImagePush
328:601:func (cli *Client) ImagePush(ctx context.Context, image string, options types.ImagePushOptions) (io.ReadCloser, error)
329:602:ImagePush requests the docker host to push an image to a remote registry. It executes the privileged function if the operation is unauthorized and it tries one more time. It's up to the caller to handle the io.ReadCloser and close it properly.
330:604:func (*Client) ImageRemove
331:605:func (cli *Client) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)
332:608:func (*Client) ImageSave
333:609:func (cli *Client) ImageSave(ctx context.Context, imageIDs []string) (io.ReadCloser, error)
334:612:func (*Client) ImageSearch
335:613:func (cli *Client) ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error)
336:616:func (*Client) ImageTag
337:617:func (cli *Client) ImageTag(ctx context.Context, source, target string) error
338:620:func (*Client) ImagesPrune
339:621:func (cli *Client) ImagesPrune(ctx context.Context, pruneFilters filters.Args) (types.ImagesPruneReport, error)
340:624:func (*Client) Info
341:625:func (cli *Client) Info(ctx context.Context) (types.Info, error)
342:628:func (*Client) NegotiateAPIVersion
343:629:func (cli *Client) NegotiateAPIVersion(ctx context.Context)
344:632:func (*Client) NegotiateAPIVersionPing
345:633:func (cli *Client) NegotiateAPIVersionPing(p types.Ping)
346:636:func (*Client) NetworkConnect
347:637:func (cli *Client) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error
348:640:func (*Client) NetworkCreate
349:641:func (cli *Client) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)
350:644:func (*Client) NetworkDisconnect
351:645:func (cli *Client) NetworkDisconnect(ctx context.Context, networkID, containerID string, force bool) error
352:648:func (*Client) NetworkInspect
353:649:func (cli *Client) NetworkInspect(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, error)
354:652:func (*Client) NetworkInspectWithRaw
355:653:func (cli *Client) NetworkInspectWithRaw(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error)
356:656:func (*Client) NetworkList
357:657:func (cli *Client) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error)
358:660:func (*Client) NetworkRemove
359:661:func (cli *Client) NetworkRemove(ctx context.Context, networkID string) error
360:664:func (*Client) NetworksPrune
361:665:func (cli *Client) NetworksPrune(ctx context.Context, pruneFilters filters.Args) (types.NetworksPruneReport, error)
362:668:func (*Client) NewVersionError
363:669:func (cli *Client) NewVersionError(APIrequired, feature string) error
364:672:func (*Client) NodeInspectWithRaw
365:673:func (cli *Client) NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error)
366:676:func (*Client) NodeList
367:677:func (cli *Client) NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error)
368:680:func (*Client) NodeRemove
369:681:func (cli *Client) NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error
370:684:func (*Client) NodeUpdate
371:685:func (cli *Client) NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error
372:688:func (*Client) Ping
373:689:func (cli *Client) Ping(ctx context.Context) (types.Ping, error)
374:692:func (*Client) PluginCreate
375:693:func (cli *Client) PluginCreate(ctx context.Context, createContext io.Reader, createOptions types.PluginCreateOptions) error
376:696:func (*Client) PluginDisable
377:697:func (cli *Client) PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error
378:700:func (*Client) PluginEnable
379:701:func (cli *Client) PluginEnable(ctx context.Context, name string, options types.PluginEnableOptions) error
380:704:func (*Client) PluginInspectWithRaw
381:705:func (cli *Client) PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error)
382:708:func (*Client) PluginInstall
383:709:func (cli *Client) PluginInstall(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
384:712:func (*Client) PluginList
385:713:func (cli *Client) PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error)
386:716:func (*Client) PluginPush
387:717:func (cli *Client) PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error)
388:720:func (*Client) PluginRemove
389:721:func (cli *Client) PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error
390:724:func (*Client) PluginSet
391:725:func (cli *Client) PluginSet(ctx context.Context, name string, args []string) error
392:728:func (*Client) PluginUpgrade
393:729:func (cli *Client) PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
394:732:func (*Client) RegistryLogin
395:733:func (cli *Client) RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)
396:736:func (*Client) SecretCreate
397:737:func (cli *Client) SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error)
398:740:func (*Client) SecretInspectWithRaw
399:741:func (cli *Client) SecretInspectWithRaw(ctx context.Context, id string) (swarm.Secret, []byte, error)
400:744:func (*Client) SecretList
401:745:func (cli *Client) SecretList(ctx context.Context, options types.SecretListOptions) ([]swarm.Secret, error)
402:748:func (*Client) SecretRemove
403:749:func (cli *Client) SecretRemove(ctx context.Context, id string) error
404:752:func (*Client) SecretUpdate
405:753:func (cli *Client) SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error
406:756:func (*Client) ServerVersion
407:757:func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error)
408:760:func (*Client) ServiceCreate
409:761:func (cli *Client) ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error)
410:764:func (*Client) ServiceInspectWithRaw
411:765:func (cli *Client) ServiceInspectWithRaw(ctx context.Context, serviceID string, opts types.ServiceInspectOptions) (swarm.Service, []byte, error)
412:768:func (*Client) ServiceList
413:769:func (cli *Client) ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
414:772:func (*Client) ServiceLogs
415:773:func (cli *Client) ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
416:777:func (*Client) ServiceRemove
417:778:func (cli *Client) ServiceRemove(ctx context.Context, serviceID string) error
418:781:func (*Client) ServiceUpdate
419:782:func (cli *Client) ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) (types.ServiceUpdateResponse, error)
420:785:func (*Client) SetCustomHTTPHeaders
421:786:func (cli *Client) SetCustomHTTPHeaders(headers map[string]string)
422:789:func (*Client) SwarmGetUnlockKey
423:790:func (cli *Client) SwarmGetUnlockKey(ctx context.Context) (types.SwarmUnlockKeyResponse, error)
424:793:func (*Client) SwarmInit
425:794:func (cli *Client) SwarmInit(ctx context.Context, req swarm.InitRequest) (string, error)
426:797:func (*Client) SwarmInspect
427:798:func (cli *Client) SwarmInspect(ctx context.Context) (swarm.Swarm, error)
428:801:func (*Client) SwarmJoin
429:802:func (cli *Client) SwarmJoin(ctx context.Context, req swarm.JoinRequest) error
430:805:func (*Client) SwarmLeave
431:806:func (cli *Client) SwarmLeave(ctx context.Context, force bool) error
432:809:func (*Client) SwarmUnlock
433:810:func (cli *Client) SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error
434:813:func (*Client) SwarmUpdate
435:814:func (cli *Client) SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error
436:817:func (*Client) TaskInspectWithRaw
437:818:func (cli *Client) TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error)
438:821:func (*Client) TaskList
439:822:func (cli *Client) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error)
440:825:func (*Client) TaskLogs
441:826:func (cli *Client) TaskLogs(ctx context.Context, taskID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
442:829:func (*Client) VolumeCreate
443:830:func (cli *Client) VolumeCreate(ctx context.Context, options volumetypes.VolumesCreateBody) (types.Volume, error)
444:833:func (*Client) VolumeInspect
445:834:func (cli *Client) VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error)
446:837:func (*Client) VolumeInspectWithRaw
447:838:func (cli *Client) VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error)
448:841:func (*Client) VolumeList
449:842:func (cli *Client) VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumesListOKBody, error)
450:845:func (*Client) VolumeRemove
451:846:func (cli *Client) VolumeRemove(ctx context.Context, volumeID string, force bool) error
452:849:func (*Client) VolumesPrune
453:850:func (cli *Client) VolumesPrune(ctx context.Context, pruneFilters filters.Args) (types.VolumesPruneReport, error)
454:853:type CommonAPIClient
455:854:type CommonAPIClient interface {
456:869:    ServerVersion(ctx context.Context) (types.Version, error)
457:871:    NegotiateAPIVersionPing(types.Ping)
458:876:type ConfigAPIClient
459:877:type ConfigAPIClient interface {
460:878:    ConfigList(ctx context.Context, options types.ConfigListOptions) ([]swarm.Config, error)
461:879:    ConfigCreate(ctx context.Context, config swarm.ConfigSpec) (types.ConfigCreateResponse, error)
462:886:type ContainerAPIClient
463:887:type ContainerAPIClient interface {
464:888:    ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
465:889:    ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
466:892:    ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error)
467:893:    ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
468:894:    ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
469:895:    ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
470:896:    ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
471:898:    ContainerInspect(ctx context.Context, container string) (types.ContainerJSON, error)
472:899:    ContainerInspectWithRaw(ctx context.Context, container string, getSize bool) (types.ContainerJSON, []byte, error)
473:901:    ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
474:902:    ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
475:904:    ContainerRemove(ctx context.Context, container string, options types.ContainerRemoveOptions) error
476:906:    ContainerResize(ctx context.Context, container string, options types.ResizeOptions) error
477:908:    ContainerStatPath(ctx context.Context, container, path string) (types.ContainerPathStat, error)
478:909:    ContainerStats(ctx context.Context, container string, stream bool) (types.ContainerStats, error)
479:910:    ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error
480:916:    CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error)
481:917:    CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error
482:918:    ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
483:922:type DistributionAPIClient
484:923:type DistributionAPIClient interface {
485:928:type ImageAPIClient
486:929:type ImageAPIClient interface {
487:930:    ImageBuild(ctx context.Context, context io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
488:931:    BuildCachePrune(ctx context.Context) (*types.BuildCachePruneReport, error)
489:932:    ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error)
490:934:    ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error)
491:935:    ImageInspectWithRaw(ctx context.Context, image string) (types.ImageInspect, []byte, error)
492:936:    ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
493:937:    ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)
494:938:    ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error)
495:939:    ImagePush(ctx context.Context, ref string, options types.ImagePushOptions) (io.ReadCloser, error)
496:940:    ImageRemove(ctx context.Context, image string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)
497:941:    ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error)
498:944:    ImagesPrune(ctx context.Context, pruneFilter filters.Args) (types.ImagesPruneReport, error)
499:948:type NetworkAPIClient
500:949:type NetworkAPIClient interface {
501:951:    NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)
502:953:    NetworkInspect(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, error)
503:954:    NetworkInspectWithRaw(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error)
504:955:    NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error)
505:957:    NetworksPrune(ctx context.Context, pruneFilter filters.Args) (types.NetworksPruneReport, error)
506:961:type NodeAPIClient
507:962:type NodeAPIClient interface {
508:964:    NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error)
509:965:    NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error
510:970:type PluginAPIClient
511:971:type PluginAPIClient interface {
512:972:    PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error)
513:973:    PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error
514:974:    PluginEnable(ctx context.Context, name string, options types.PluginEnableOptions) error
515:975:    PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error
516:976:    PluginInstall(ctx context.Context, name string, options types.PluginInstallOptions) (io.ReadCloser, error)
517:977:    PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (io.ReadCloser, error)
518:980:    PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error)
519:981:    PluginCreate(ctx context.Context, createContext io.Reader, options types.PluginCreateOptions) error
520:985:type SecretAPIClient
521:986:type SecretAPIClient interface {
522:987:    SecretList(ctx context.Context, options types.SecretListOptions) ([]swarm.Secret, error)
523:988:    SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error)
524:995:type ServiceAPIClient
525:996:type ServiceAPIClient interface {
526:997:    ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error)
527:998:    ServiceInspectWithRaw(ctx context.Context, serviceID string, options types.ServiceInspectOptions) (swarm.Service, []byte, error)
528:999:    ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
529:1001:    ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) (types.ServiceUpdateResponse, error)
530:1002:    ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
531:1003:    TaskLogs(ctx context.Context, taskID string, options types.ContainerLogsOptions) (io.ReadCloser, error)
532:1005:    TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error)
533:1009:type SwarmAPIClient
534:1010:type SwarmAPIClient interface {
535:1013:    SwarmGetUnlockKey(ctx context.Context) (types.SwarmUnlockKeyResponse, error)
536:1021:type SystemAPIClient
537:1022:type SystemAPIClient interface {
538:1023:    Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
539:1024:    Info(ctx context.Context) (types.Info, error)
540:1025:    RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)
541:1026:    DiskUsage(ctx context.Context) (types.DiskUsage, error)
542:1027:    Ping(ctx context.Context) (types.Ping, error)
543:1031:type VolumeAPIClient
544:1032:type VolumeAPIClient interface {
545:1033:    VolumeCreate(ctx context.Context, options volumetypes.VolumesCreateBody) (types.Volume, error)
546:1034:    VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error)
547:1035:    VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error)
548:1036:    VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumesListOKBody, error)
549:1038:    VolumesPrune(ctx context.Context, pruneFilter filters.Args) (types.VolumesPruneReport, error)
550:1045:api/types	Package types is used for API stability in the types and response to the consumers of the API stats endpoint.
554:1049:api/types/filters	Package filters provides tools for encoding a mapping of keys to a set of multiple values.
561:1056:api/types/swarm/runtime	Package runtime is a generated protocol buffer package.
564:1059:api/types/versions/v1p19	Package v1p19 provides specific API types for the API version 1, patch 19.
565:1060:api/types/versions/v1p20	Package v1p20 provides specific API types for the API version 1, patch 20.
