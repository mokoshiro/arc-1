workspace(
    name = "arc",
)

load(
    "@bazel_tools//tools/build_defs/repo:http.bzl",
    "http_archive",
)

### Rules_go and gazelle
http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.19.1/rules_go-0.19.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.1/rules_go-0.19.1.tar.gz",
    ],
    sha256 = "8df59f11fb697743cbb3f26cfb8750395f30471e9eabde0d174c3aebc7a1cd39",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load(
    "@bazel_gazelle//:deps.bzl",
    "go_repository",
)

# Override rules_go dependencies
go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    tag = "v1.24.0",
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_rules_dependencies",
    "go_register_toolchains",
)

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.11.5",
)

load(
    "@bazel_gazelle//:deps.bzl",
    "gazelle_dependencies",
)

gazelle_dependencies()

### Docker
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "aed1c249d4ec8f703edddf35cbe9dfaca0b5f5ea6e4cd9e83e99f3b0d1136c3d",
    strip_prefix = "rules_docker-0.7.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.7.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

load(
    "@bazel_tools//tools/build_defs/repo:git.bzl",
    "new_git_repository",
    "git_repository",
)
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

### Proto dependencies
git_repository(
    name = "com_google_protobuf",
    commit = "09745575a923640154bcf307fba8aedff47f240a",
    remote = "https://github.com/protocolbuffers/protobuf",
    shallow_since = "1558721209 -0700",
)

load(
    "@com_google_protobuf//:protobuf_deps.bzl",
    "protobuf_deps",
)

protobuf_deps()

http_archive(
    name = "build_stack_rules_proto",
    urls = ["https://github.com/stackb/rules_proto/archive/1d6550fc2e62.tar.gz"],
    sha256 = "113e6792f5b20679285c86d91c163cc8c4d2b4d24d7a087ae4f233b5d9311012",
    strip_prefix = "rules_proto-1d6550fc2e625d47dc4faadac92d7cb20e3ba5c5",
)

### CGO libraries
new_git_repository(
    name = "com_github_uber_h3_go",
    remote = "https://github.com/uber/h3-go.git",
    commit = "ba79e9fc50a2cbf5ab898f6268dc507f1037ef26",
    build_file = "@arc//:builds/h3.BUILD",
)

### Auto generated

go_repository(
    name = "com_github_armon_consul_api",
    commit = "eb2c6b5be1b6",
    importpath = "github.com/armon/consul-api",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    tag = "v0.3.1",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    tag = "v3.3.10",
)

go_repository(
    name = "com_github_coreos_go_etcd",
    importpath = "github.com/coreos/go-etcd",
    tag = "v2.0.0",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    tag = "v0.2.0",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man",
    importpath = "github.com/cpuguy83/go-md2man",
    tag = "v1.0.10",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    tag = "v1.4.7",
)

go_repository(
    name = "com_github_garyburd_redigo",
    importpath = "github.com/garyburd/redigo",
    tag = "v1.6.0",
)

go_repository(
    name = "com_github_gin_contrib_sse",
    commit = "5545eab6dad3",
    importpath = "github.com/gin-contrib/sse",
)

go_repository(
    name = "com_github_gin_gonic_gin",
    importpath = "github.com/gin-gonic/gin",
    tag = "v1.4.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "v1.3.2",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    tag = "v1.1.6",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    tag = "v1.8.0",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    tag = "v0.0.8",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    importpath = "github.com/mitchellh/go-homedir",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    tag = "v1.1.2",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    commit = "bacd9c7ef1dd",
    importpath = "github.com/modern-go/concurrent",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_russross_blackfriday",
    importpath = "github.com/russross/blackfriday",
    tag = "v1.5.2",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    tag = "v1.1.2",
)

go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    tag = "v0.0.5",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    tag = "v1.0.3",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    tag = "v1.3.2",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    tag = "v0.1.1",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    tag = "v1.4.0",
)

go_repository(
    name = "com_github_ugorji_go",
    importpath = "github.com/ugorji/go",
    tag = "v1.1.4",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    commit = "d75b2dcb6bc8",
    importpath = "github.com/ugorji/go/codec",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    commit = "b2862e3d0a77",
    importpath = "github.com/xordataexchange/crypt",
)

go_repository(
    name = "in_gopkg_check_v1",
    commit = "41f04d3bba15",
    importpath = "gopkg.in/check.v1",
)

go_repository(
    name = "in_gopkg_go_playground_assert_v1",
    importpath = "gopkg.in/go-playground/assert.v1",
    tag = "v1.2.1",
)

go_repository(
    name = "in_gopkg_go_playground_validator_v8",
    importpath = "gopkg.in/go-playground/validator.v8",
    tag = "v8.18.2",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    tag = "v2.2.4",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "c2843e01d9a2",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "3b0461eec859",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "3ef323f4f1fd",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    tag = "v0.3.2",
)

go_repository(
    name = "co_honnef_go_tools",
    commit = "ea95bdfd59fc",
    importpath = "honnef.co/go/tools",
)

go_repository(
    name = "com_github_alicebob_gopher_json",
    commit = "5a6b3ba71ee6",
    importpath = "github.com/alicebob/gopher-json",
)

go_repository(
    name = "com_github_alicebob_miniredis_v2",
    importpath = "github.com/alicebob/miniredis/v2",
    tag = "v2.9.0",
)

go_repository(
    name = "com_github_apmckinlay_gsuneido",
    commit = "1f10244968e3",
    importpath = "github.com/apmckinlay/gsuneido",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    tag = "v1.21.8",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    commit = "27f122750802",
    importpath = "github.com/BurntSushi/xgb",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    tag = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    commit = "2972be24d48e",
    importpath = "github.com/chzyer/readline",
)

go_repository(
    name = "com_github_chzyer_test",
    commit = "a1ea475d72b1",
    importpath = "github.com/chzyer/test",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    tag = "v0.3.4",
)

go_repository(
    name = "com_github_cockroachdb_apd",
    importpath = "github.com/cockroachdb/apd",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_data_dog_go_sqlmock",
    importpath = "github.com/DATA-DOG/go-sqlmock",
    tag = "v1.3.3",
)

go_repository(
    name = "com_github_davidbyttow_govips",
    commit = "d272f04c0fea",
    importpath = "github.com/davidbyttow/govips",
)

go_repository(
    name = "com_github_ericlagergren_decimal",
    commit = "73749d4874d5",
    importpath = "github.com/ericlagergren/decimal",
)

go_repository(
    name = "com_github_go_redsync_redsync",
    importpath = "github.com/go-redsync/redsync",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_gofrs_uuid",
    importpath = "github.com/gofrs/uuid",
    tag = "v3.2.0",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    tag = "v1.2.1",
)

go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b",
    importpath = "github.com/golang/glog",
)

go_repository(
    name = "com_github_golang_groupcache",
    commit = "869f871628b6",
    importpath = "github.com/golang/groupcache",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_gomodule_redigo",
    importpath = "github.com/gomodule/redigo",
    tag = "v2.0.0",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    tag = "v0.3.0",
)

go_repository(
    name = "com_github_google_go_github_v27",
    importpath = "github.com/google/go-github/v27",
    tag = "v27.0.4",
)

go_repository(
    name = "com_github_google_go_querystring",
    importpath = "github.com/google/go-querystring",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    tag = "v2.1.0",
)

go_repository(
    name = "com_github_google_pprof",
    commit = "54271f7e092f",
    importpath = "github.com/google/pprof",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    tag = "v2.0.5",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    tag = "v0.5.1",
)

go_repository(
    name = "com_github_hayabusa_havana",
    importpath = "github.com/hayabusa/havana",
    tag = "v0.0.9",
)

go_repository(
    name = "com_github_jeffail_gabs",
    importpath = "github.com/Jeffail/gabs",
    tag = "v1.4.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    commit = "c2b33e8439af",
    importpath = "github.com/jmespath/go-jmespath",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    commit = "af01ea7f8024",
    importpath = "github.com/jstemmer/go-junit-report",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_knetic_govaluate",
    importpath = "github.com/Knetic/govaluate",
    tag = "v3.0.0",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_lukechampine_freeze",
    commit = "f514e08ae5a0",
    importpath = "github.com/lukechampine/freeze",
)

go_repository(
    name = "com_github_mattn_go_scan",
    commit = "c32d62d79baf",
    importpath = "github.com/mattn/go-scan",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    tag = "v0.8.1",
)

go_repository(
    name = "com_github_rs_xid",
    importpath = "github.com/rs/xid",
    tag = "v1.2.1",
)

go_repository(
    name = "com_github_rwcarlsen_goexif",
    commit = "9e8deecbddbd",
    importpath = "github.com/rwcarlsen/goexif",
)

go_repository(
    name = "com_github_shopspring_decimal",
    commit = "cd690d0c9e24",
    importpath = "github.com/shopspring/decimal",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    tag = "v1.4.2",
)

go_repository(
    name = "com_github_stvp_tempredis",
    commit = "b82af8480203",
    importpath = "github.com/stvp/tempredis",
)

go_repository(
    name = "com_github_ua_parser_uap_go",
    commit = "ac4b80194e5b",
    importpath = "github.com/ua-parser/uap-go",
)

go_repository(
    name = "com_github_volatiletech_inflect",
    commit = "e7201282ae8d",
    importpath = "github.com/volatiletech/inflect",
)

go_repository(
    name = "com_github_volatiletech_null",
    importpath = "github.com/volatiletech/null",
    tag = "v8.0.0",
)

go_repository(
    name = "com_github_volatiletech_sqlboiler",
    importpath = "github.com/volatiletech/sqlboiler",
    tag = "v3.4.0",
)

go_repository(
    name = "com_github_yuin_gopher_lua",
    commit = "1cd887cd7036",
    importpath = "github.com/yuin/gopher-lua",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    tag = "v0.26.0",
)

go_repository(
    name = "in_gopkg_go_playground_colors_v1",
    importpath = "gopkg.in/go-playground/colors.v1",
    tag = "v1.2.0",
)

go_repository(
    name = "in_gopkg_gographics_imagick_v3",
    importpath = "gopkg.in/gographics/imagick.v3",
    tag = "v3.2.0",
)

go_repository(
    name = "in_gopkg_h2non_bimg_v1",
    importpath = "gopkg.in/h2non/bimg.v1",
    tag = "v1.0.19",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    tag = "v0.9.1",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    tag = "v0.22.1",
)

go_repository(
    name = "io_rsc_binaryregexp",
    importpath = "rsc.io/binaryregexp",
    tag = "v0.2.0",
)

go_repository(
    name = "org_go4",
    commit = "94abd6928b1d",
    importpath = "go4.org",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    tag = "v0.7.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    tag = "v1.4.0",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "24fa4b261c55",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_x_exp",
    commit = "509febef88a4",
    importpath = "golang.org/x/exp",
)

go_repository(
    name = "org_golang_x_image",
    commit = "1bd0cf576493",
    importpath = "golang.org/x/image",
)

go_repository(
    name = "org_golang_x_lint",
    commit = "d0100b6bd8b3",
    importpath = "golang.org/x/lint",
)

go_repository(
    name = "org_golang_x_mobile",
    commit = "d3739f865fa6",
    importpath = "golang.org/x/mobile",
)

go_repository(
    name = "org_golang_x_oauth2",
    commit = "d2e6202438be",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "112230192c58",
    importpath = "golang.org/x/sync",
)

go_repository(
    name = "org_golang_x_time",
    commit = "9d24e82272b4",
    importpath = "golang.org/x/time",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "2c0ae7006135",
    importpath = "golang.org/x/tools",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    tag = "v1.4.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    tag = "v1.1.0",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    tag = "v1.10.0",
)

go_repository(
    name = "org_golang_x_xerrors",
    commit = "1b5146add898",
    importpath = "golang.org/x/xerrors",
)

go_repository(
    name = "com_github_uber_h3_go",
    importpath = "github.com/uber/h3-go",
    tag = "v3.0.1",
)

go_repository(
    name = "com_github_k0kubun_colorstring",
    commit = "9440f1994b88",
    importpath = "github.com/k0kubun/colorstring",
)

go_repository(
    name = "com_github_k0kubun_pp",
    importpath = "github.com/k0kubun/pp",
    tag = "v3.0.1",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    tag = "v0.1.4",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    tag = "v1.4.1",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    tag = "v1.4.1",
)
