build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main max.go
deploy_prod: build
	serverless deploy --stage prod
destroy:
	serverless remove --stage prod

# Line 3 
# 	serverless deploy --stage prod --aws-profile myawsprofileFile
# 	(if profile not specify = default to main cred) 
# Line 5 
# 	(added this one in myself because I had trouble deleting the resources)
