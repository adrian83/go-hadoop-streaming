
HERE=$(pwd)

echo -e "Checking environment variables..."


if [ -z "$JAVA_HOME" ]; then
  echo "Variable JAVA_HOME not set!"
  echo "Set JAVA_HOME. Example: 'export JAVA_HOME=/path/to/jdk_or_jre'"
  echo "Aborting!"
  exit 1
else
  echo "JAVA_HOME="$JAVA_HOME
fi

if [ -z "$HADOOP_HOME" ]; then
  echo "Variable HADOOP_HOME not set!"
  echo "Set HADOOP_HOME. Example: 'export HADOOP_HOME=/path/to/hadoop'"
  echo "Aborting!"
  exit 1
else
  echo "HADOOP_HOME="$HADOOP_HOME
fi

last_dir=$(basename "$HADOOP_HOME")
HADOOP_VERSION=$(echo "$last_dir" | sed -E 's/.*-([0-9]+\.[0-9]+\.[0-9]+).*/\1/')
echo "HADOOP_VERSION="$HADOOP_VERSION


echo -e "Checking installed software..."

GO=${GO:-$(which go)}
if [ -z "$GO" ]; then
  echo "Go not installed. Please install Go"
  echo "Aborting!"
  exit 1
else
  echo "GO="$GO
fi


echo -e "Building Mapper and Reducer..."

cd $HERE/cmd/mapper && go build
cd $HERE/cmd/reducer && go build
cd $HERE

echo -e "Running Hadoop job..."
now=${now:-$(date +"%Y_%d_%d_%H_%M_%S")}
OUTPUT="./output_$now"

echo "$HADOOP_HOME/bin/hadoop  jar $HADOOP_HOME/share/hadoop/tools/lib/hadoop-streaming-$HADOOP_VERSION.jar -input ./input -output $OUTPUT -mapper ./mapper/mapper -reducer ./reducer/reducer"
$HADOOP_HOME/bin/hadoop  jar $HADOOP_HOME/share/hadoop/tools/lib/hadoop-streaming-$HADOOP_VERSION.jar -input ./input -output $OUTPUT -mapper ./cmd/mapper/mapper -reducer ./cmd/reducer/reducer

echo "Check output directory ($OUTPUT) to see the results"
