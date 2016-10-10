

kubectl run huoneisto-utils --image=gcr.io/jntlserv0/huoneisto_utils:0.2 --env="THEMES=realestate" --env="LOCALE=fi_FI" --env=""  --env=""  --env=""  --env=""  --env=""



kubectl delete pods $( kubectl get pods --show-all |grep 'Completed' |awk '{print $1}')

kubectl delete pods $( kubectl get pods --show-all |grep 'Completed' |awk '{print $1}') && kubectl delete jobs pi && kubectl create -f job_espoo.huoneisto.mobi.yaml && kubectl create -f job_www.huoneisto.mobi.yaml && kubectl create -f job_huoneisto.mobi.yaml
