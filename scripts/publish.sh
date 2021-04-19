ARCH=$INPUT_ARCHITECTURE
DISTRO=""
OS=$INPUT_OS

if [[ "$OS" = *"ubuntu"* ]]; then
  DISTRO="ubuntu"
elif [[ "$OS" = *"macos"* ]]; then
  DISTRO="darwin"
else
  printf "ERROR\tOperating system %s not supported\n" "$OS"
  exit 1
fi

if [ -z "$DISTRO" ]; then
  printf "ERROR\tUnable to detect the operating system\n"
  exit 1
fi

printf "INFO\tBAZEL FOLDER\n"
bazel info -c opt bazel-bin

ROOT_FOLDER=$(bazel info -c opt bazel-bin)

# Optional personal access token for external repository
TOKEN=$GITHUB_TOKEN
if ! [[ -z ${INPUT_TOKEN} ]]; then
  TOKEN=$INPUT_TOKEN
fi

if ! [[ -f "$CLIENT_BINARY" ]]; then
    printf "$CLIENT_BINARY does not exist"
fi

if ! [[ -f "$SERVICE_BINARY" ]]; then
    printf "$SERVICE_BINARY does not exist"
fi

if ! [[ -f "$TEST_SERVER_BINARY" ]]; then
    printf "$TEST_SERVER_BINARY does not exist"
fi

if ! [[ -f "$OUTPUT_TRANSFORM_BINARY" ]]; then
    printf "$OUTPUT_TRANSFORM_BINARY does not exist"
fi

# bundle the binaries
if ! type "tar" > /dev/null 2>&1; then
  printf "ERROR\ttar not found\n"
  exit 1;
fi

ASSET_SUFFIX="$DISTRO-$ARCH-$INPUT_VERSION.tar.gz"
for binary in nighthawk_client nighthawk_service nighthawk_test_server nighthawk_output_transform
do
    if ! tar -zcvf $binary-$ASSET_SUFFIX --directory=$ROOT_FOLDER $binary; then
        printf "ERROR\tUnable to create bundle\n"
    fi
    # Upload artifact
    GITHUB_API_URL="api.github.com"
    RELEASE_URL="https://$GITHUB_API_URL/repos/$INPUT_REPO/releases"
    RELEASE_UPLOAD_URL=$(curl -H "Authorization: token $TOKEN" $RELEASE_URL | jq -r '.[] | select(.tag_name == "'${INPUT_VERSION}'")' | jq -r .upload_url)
    pattern="{?"
    RELEASE_ASSET_URL="${RELEASE_UPLOAD_URL%$pattern*}"
    response=$(curl -X POST -H "Content-Type: application/tar" -H "Authorization: token $TOKEN" -H "Accept: application/vnd.github.v3+json" --data-binary @$binary-$ASSET_SUFFIX "$RELEASE_ASSET_URL?name=$binary-$ASSET_SUFFIX")
    if ! response; then
        printf "ERROR\tUnable to post the artifact\n"
    fi
    printf "INFO\tUploaded $binary succesfully!!\n"
done
printf "INFO\tAction ran succesfully!!\n"