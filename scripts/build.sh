cd ../nighthawk

if ! bazel build -c opt //:nighthawk_client; then
  printf "ERROR\tUnable to build nighthawk client\n"
  exit 1;
fi

if ! bazel build -c opt //:nighthawk_service; then
  printf "ERROR\tUnable to build nighthawk service\n"
  exit 1;
fi

if ! bazel build -c opt //:nighthawk_test_server; then
  printf "ERROR\tUnable to build nighthawk test server\n"
  exit 1;
fi