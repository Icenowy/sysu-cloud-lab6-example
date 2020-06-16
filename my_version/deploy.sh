scp my_deployment.yaml my_service.yaml ice-lab5:/home/icenowy/
ssh ice-lab5 kubectl apply -f /home/icenowy/my_deployment.yaml
ssh ice-lab5 kubectl apply -f /home/icenowy/my_service.yaml
