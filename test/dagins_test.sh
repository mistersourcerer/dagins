set -e

# ensure there is a bin/dagins first?
./build.sh

usage=$(./bin/dagins goal 2>&1)
failures=()

if [[ ${usage} != *"Usage of goal"* ]]; then
  failure="FAIL: ${0}:${LINENO}: show usage"
  failures+=($failure)
fi


if [[ ${#failures[*]} -eq 0 ]]; then
  echo "all passed!"
  exit
else
  for failure in "${failures[*]}"; do
    echo $failure
  done

  exit -1
fi
