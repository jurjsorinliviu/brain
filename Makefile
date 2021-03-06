.PHONY: install

install: 
	cd examples/training && glide install
	cd examples/inference && glide install

test: test_bnn

test_bnn: datasets/mnist_train.csv datasets/mnist_test.csv
	@( go test ./bnn/ )

goget:
	@( \
		go get github.com/ReconfigureIO/brain;
	)

gogetu:
	@( \
		go get -u github.com/ReconfigureIO/brain;
	)

datasets/mnist_train.csv:
	unzip datasets/mnist_train.csv.zip -d datasets/

datasets/mnist_test.csv:
	unzip datasets/mnist_test.csv.zip -d datasets/