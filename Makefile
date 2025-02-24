
# make i c=コンテストID t=問題ID
i:
	./shell/init.sh $(c) $(t) $(diff)

r:
	./shell/run.sh

t:
	./shell/test.sh

s:
	./shell/submit.sh

url:
	./shell/submiturl.sh

login:
	oj login https://atcoder.jp
