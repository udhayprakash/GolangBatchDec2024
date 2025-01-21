# AWS Lambda

## AWs Lamdba Execution environment lifecycle

    https://docs.aws.amazon.com/images/lambda/latest/dg/images/Overview-Full-Sequence.png

    1) Init phase
        - lambda creates or unfreezes execution environment with configured reosurces 
            - downloads code for function 
            - all layers 
            - initializes any extensions
            - initializes the runtime
            - runs functions's initialization code (code outside main handler)
            - Init stage happens for first function, or in advance if enabled provisioned concurrency
            - Init Phase: Extension init, Runtime init, Function init 
    2) Invoke phase 
        - lambda invoked function handler
        - then other functions, based on logic

    3) shutdown phase
        - triggered if lamdba function does not receive any invocatoon for a period of time. 
        - shuts down runtime, altersts the extensions to let them stop cleanly, removes envioronemnt

## Usecases

    1) FileProcessing, with s3 bucket
    2) Stream Processing (lambda, Kinesis)
    3) Web Application
