build:
	@docker build -t go-skeleton .

run:
	@docker run -v `pwd`/data:/data go-skeleton
