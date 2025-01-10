# Mocking


    go get github.com/golang/mock/gomock
    <!-- go get github.com/golang/mock/mockgen -->
        go install github.com/golang/mock/mockgen@latest


mockgen -source=calculator.go -destination=mock_calculator.go -package=main
