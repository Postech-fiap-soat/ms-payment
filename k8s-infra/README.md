# terraform-k8s

Requisitos:

Fazer download do aws cli: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
 - Apos instalado, executar no terminal:
    aws configure
 - Adicionar credenciais geradas ao criar um usuario IAM no aws console.

Fazer download do Terraform: https://developer.hashicorp.com/terraform/downloads?product_intent=terraform
Fazer download do aws-iam-authenticator: https://docs.aws.amazon.com/pt_br/eks/latest/userguide/install-aws-iam-authenticator.html
Fazer download do kubectl: https://kubernetes.io/docs/tasks/tools/
Executar: 
- terraform init
- terraform apply --auto-approve 
- cp kubeconfig ~/.kube/config
- bash /k8s/deploy-eks.sh
