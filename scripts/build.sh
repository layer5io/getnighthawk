NIGHTHAWK_DIR="/home/runner/work/getnighthawk/getnighthawk/"
cd $NIGHTHAWK_DIR

if ! bazel build -c opt --define tcmalloc=gperftools //:nighthawk; then
  printf "ERROR\tUnable to build nighthawk client\n"
  exit 1;
fi
