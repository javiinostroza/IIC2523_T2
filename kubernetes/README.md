# Crear un Cluster

Para crear un closter se debe utilizar el comando: minikube start
este comando iniciara un cluster,
para obtener información del cluster se debe utilizar el comando: kubectl cluster-info
con el siguiente comando se aprendera sobre sus nodos: kubectl get nodes

# Desplegar un cluster

para despegar un cluster se debe usar el siguiente comando, entregandole el nombre
del cluster y el lugar donde esta la imagen: 
kubectl create deployment nombre --image=lugar

Para poder ver los cluster desplegados se ocupa este comando: kubectl get deployments

# Troubleshoot Kubernetes

Para obtener los pods se utiliza: kubectl get pods

Para obtener las propiedades como la IP, el puerto y eventos de los pods: kubectl describe pods

Obtener los logs de un pod con su nombre: kubectl logs nombre

con curl se puede ser si el cluster sta corriedo: curl http://localhost:8001

Con este comando se pueden correr programas en el cluster: kubectl exec $POD_NAME

# Exponer una aplicación públicamente

Para exponer un cluster se debe usar el siguiente comando entregando el nombre
del cluster y el puerto a ser expuesto:

kubectl expose deployment/nombre --type="NodePort" --port 8080

Para borrar un servicio se ocupa el siguiente comando con el nombre del cluster:

kubectl delete service -l app=nombre


# Escalar una aplicación

Para escalar un cluster se debe usar el siguiente comando,entregando el nombre
y el numero de replicas a poner (esto sirve para aumentar como disminuir):

kubectl scale deployments/nombre --replicas=numero

# Actualizar una aplicación.

Para actualizar un cluster se debe usar el siguiente comando, entregando el nombre del cluster
seguido por el nombre e igualandolo a la nueva imagen:

kubectl set image deployments/nombre nombre=jocatalin/nombre_img


