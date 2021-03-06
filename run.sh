
HERE=$(pwd)

echo -e "\n  CHECKING ENV VARS \n"

if [ -z "$JAVA_HOME" ]; then
    echo "    Please set JAVA_HOME"
    exit 1
else
  echo "    JAVA_HOME=" $JAVA_HOME
fi

if [ -z "$HADOOP_HOME" ]; then
    echo "    Please set HADOOP_HOME"
    exit 1
else
  echo "    HADOOP_HOME=" $HADOOP_HOME
fi

if [ -z "$HADOOP_VERSION" ]; then
    echo "    Please set HADOOP_VERSION"
    exit 1
else
  echo "    HADOOP_VERSION=" $HADOOP_VERSION
fi



echo -e "\n  CHECKING INSTALLED SOFTWARE \n"

GO=${GO:-$(which go)}

if [ -z "$GO" ]; then
    echo "    Please install Go (golang)"
    exit 1
else
  echo "    GO=" $GO
fi


echo -e "\n  BUILD MAPPER AND REDUCER \n"

cd $HERE/cmd/mapper && go build
cd $HERE/cmd/reducer && go build
cd $HERE


echo -e "\n  RUN HADOOP JOB \n"

OUTPUT="./output"

if [ -d "$OUTPUT" ]; then
  echo "    Please delete $OUTPUT directory"
  exit 1
else
  echo "    $HADOOP_HOME/bin/hadoop  jar $HADOOP_HOME/share/hadoop/tools/lib/hadoop-streaming-$HADOOP_VERSION.jar -input ./input -output ./output -mapper ./mapper/mapper -reducer ./reducer/reducer"

  $HADOOP_HOME/bin/hadoop  jar $HADOOP_HOME/share/hadoop/tools/lib/hadoop-streaming-$HADOOP_VERSION.jar -input ./input -output ./output -mapper ./cmd/mapper/mapper -reducer ./cmd/reducer/reducer

  echo "    Check output directory to see the results"
fi
