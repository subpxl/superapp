argocd 

needs podman , kubectl, kind

cmds
kubectl cluster-info

kubectl create namespace argocd
kubectl config set-context --current --namespace=argocd

kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server
kubectl port-forward svc/argocd-server -n argocd 8080:443
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo


to setup


Step 5: Create Your First Application
Click on the "New App" button in the Argo CD UI.

Fill in the application details:

Application Name: A unique name for your application.
Project: Default.
Sync Policy: Manual (you can change this later to automatic if desired).
Repository URL: The Git repository containing your Kubernetes manifests.
Revision: The branch, tag, or commit to use.
Path: The path within the repository where the manifests are located.
Cluster URL: Usually https://kubernetes.default.svc.
Namespace: The namespace where the application will be deployed.


Step 6: Sync the Application


to run

needa git branch with k8s manifest folder and and namespace point to it , then sync 

to view

kubectl config set-context --current --namespace=superapp
kubectl get pods
