# go-rest-sample
GoでRESTAPIサーバの勉強

### 課題概要
https://gist.github.com/koudaiii/62a9971625c9b6d4026da60f4b79dd03


### 課題

* 公開URL（課題１,２）  
https://go-rest-sample-238407.appspot.com/  
http://34.85.73.127  

* 課題３
```
docker-compose -f docker-compose-3.yml build --no-cache  
docker-compose -f docker-compose-3.yml up -d  
```
### メモ
dockerイメージ作成
```
export PROJECT_ID="$(gcloud config get-value project -q)"
docker build -t gcr.io/${PROJECT_ID}/hello-app:v1 .
gcloud docker -- push gcr.io/${PROJECT_ID}/hello-app:v1
```
gkeにデプロイ
```
kubectl run hello-web --image=gcr.io/${PROJECT_ID}/hello-app:v1 --port 8080
kubectl expose deployment hello-web --type=LoadBalancer --port 80 --target-port 8080
kubectl get service
```

参考：  
https://qiita.com/i35_267/items/274206f50c9dec980a0c  
https://www.topgate.co.jp/gcp07-how-to-start-docker-image-gke#google-container-engine-gke  
