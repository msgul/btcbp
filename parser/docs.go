package parser

/*
src: https://en.bitcoin.it/wiki/Protocol_documentation

	[Block Parts]                    [Size]
	├ magic bytes                   - 4 bytes
	├ Block Size                    - 4 bytes
	├ block header                  - 80 bytes
	│   └┬ version                  - 4 bytes
	│    ├ hash previous block      - 32 bytes
	│    ├ hash merkle root         - 32 bytes
	│    ├ time                     - 4 bytes
	│    ├ bits                     - 4 bytes
	│    └ nonce                    - 4 bytes
	├ for tx count                  - var_int
		├ tx data                  - remainder
		│	 └┬ version             - 4 bytes
		│	  ├ for input count     - var_int
		│	  │    ├ tx out hash    - 32 bytes
		│	  │    ├ tx out index   - 4 bytes
		│	  │    ├ script length  - var_int
		│	  │    └ sigscript      - script length
		│	  ├ for output count    - var_int
		│	  │    ├ value          - 8 bytes
		│	  │    ├ script length  - var_int
		│	  │    └ pkscript       - script length
		│	  └ locktime            - 4 bytes
*/
