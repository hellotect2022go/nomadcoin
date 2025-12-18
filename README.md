**Go 언어 기반 프로젝트 학습 항목**

* **1. 파일 구조 및 의존성 관리**
    * `go.mod` : `package.json` 같은거, `requirements.txt`

* **2. Blockchain & Block 의 개념**
    * 1 blockchain & block 의 개념

* **3. [GO] Server Side Rendering Web Service**
    * - template 라이브러리 , handler, router
    * - json marshal, encoder
    * - TextMarshaler 인터페이스로 MarshalText 구현
    * - multiplexer (Mux, gorilla mux)
    * - error.New 로 새 에러메세지 만들기
    * - Middlewares 설정으로 application/json 미들웨어만들기

* **4. CLI Tool 만들기**
    * - flag 라이브러리 사용

* **5. [Go] BoltDB**
    * - blockchain 의 persistence 속성부여
    * - key-value DB 사용
    * - window 에서 boltbrowser 설치 방법 
    go install github.com/br0xen/boltbrowser@latest      
    boltbrowser blockChain.db
    * - boltweb 설치 및 실행방법
    go install github.com/evnix/boltdbweb@latest
    boltdbweb --db-name=blockChain.db --port=8888
    boltdbweb --db-name=<DBfilename>[required] --port=<port>[optional] --static-path=<static-path>[optional]

    sdfa