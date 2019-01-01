## tendiesticks is a application that looks at the hottest trades of the day in r/wallstreetbets and tries to emulate it.


```
go run cmd/main.go
```

I am also hosting some/all? of the data on elasticsearch

```
sudo chkconfig --add elasticsearch
sudo -i service elasticsearch start
sudo -i service elasticsearch stop

```

check if elastic search is running
```
curl -X GET "localhost:9200/"
```

elastic logs

```
/var/log/elasticsearch/
```

kibana

```
sudo chkconfig --add kibana
sudo -i service kibana start
sudo -i service kibana stop
http://localhost:5601/app/kibana
```

```
pip3 install -r requirements.txt

```



TODO:
* Figure how to increase requests rate for Reddit
* Create polling to Alpha Vantage
* Create parse for Reddit posts
* create db to store s&p 500 financial data

Alpha Vantage API: https://www.alphavantage.co/documentation/

I'm also new to Go, so there will be use links and resources below:
https://github.com/golang-standards/project-layout

Useful Links for Go projects
https://outcrawl.com/go-elastic-search-service
Guide for elasticSearch (RPM based)
https://www.elastic.co/guide/en/elasticsearch/reference/current/rpm.html

Visualize with Kibana
https://www.elastic.co/guide/en/kibana/current/rpm.html

Tensorflow with golang
https://godoc.org/github.com/tensorflow/tensorflow/tensorflow/go

