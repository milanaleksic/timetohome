#/bin/sh

if [ -z "$GOPATH" ]; then
    echo GOPATH environment variable not set
    exit
fi

if [ ! -e "$GOPATH/bin/2goarray" ]; then
    echo "Installing 2goarray..."
    go get github.com/cratonica/2goarray
    if [ $? -ne 0 ]; then
        echo Failure executing go get github.com/cratonica/2goarray
        exit
    fi
fi

OUTPUT=iconunix.go
echo Generating $OUTPUT
echo "//+build linux darwin" > $OUTPUT
echo >> $OUTPUT
cat "home.png" | $GOPATH/bin/2goarray Data main >> $OUTPUT
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT
    exit
fi
echo Finished
