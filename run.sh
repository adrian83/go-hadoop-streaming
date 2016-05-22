
echo "----------------------------------------------------"
echo "  CHECKING ENV VARS"
echo ""





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


echo ""
echo "----------------------------------------------------"
echo "  CHECKING INSTALLED SOFTWARE"
echo ""

GO=${GO:-$(which go)}

if [ -z "$GO" ]; then
    echo "    Please install Go (golang)"
    echo ""
    echo "----------------------------------------------------"

    exit 1
else
  echo "    GO=" $GO
fi

echo ""
echo "----------------------------------------------------"

echo "  MODIFY GOPATH"
echo ""
HERE=$(pwd)
PROJ="/go_mapreduce"
GOPATH=$HERE$PROJ
echo "    GOPATH=" $GOPATH
echo ""
echo "----------------------------------------------------"

echo "  BUILD MAPPER AND REDUCER"
echo ""

go install mapreduce/mapper/
go install mapreduce/reducer/


echo ""
echo "----------------------------------------------------"

echo "  RUN HADOOP JOB"
echo ""

OUTPUT="./output"

if [ -d "$OUTPUT" ]; then
  echo "    Please delete $OUTPUT directory"

else
  echo "    $HADOOP_HOME/bin/hadoop  jar $HADOOP_HOME/share/hadoop/tools/lib/hadoop-streaming-$HADOOP_VERSION.jar -input ./input -output ./output -mapper ./go_mapreduce/bin/mapper -reducer ./go_mapreduce/bin/reducer"

  $HADOOP_HOME/bin/hadoop  jar $HADOOP_HOME/share/hadoop/tools/lib/hadoop-streaming-$HADOOP_VERSION.jar -input ./input -output ./output -mapper ./go_mapreduce/bin/mapper -reducer ./go_mapreduce/bin/reducer

  echo "    Check output directory to see the results"

fi

echo ""
echo "----------------------------------------------------"
