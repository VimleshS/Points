 
EC2 Elastic compute cloud. (virtual server in cloud)
	Auto Scaling. (part of EC2)
		-> integrates with load balancer
	Lambda (run code in response to event)
		-> we only need to provide code.
	EC2 Container services (Runs & manage docker containers)
		-> integrates with load balancer
		-> integrates with EBS
Storage and Content Delivery   (just like hard drive)
	S3 (Scalable Storage Services)
	Glacier (Archive storage in cloud)
		-> S3's cousin
		-> there is a delay
		-> like a tape drive
	EBS (Elastic block storage)
		-> persistence off-instance storage
		-> SSD(solid state drive) or mangetic disk, encryption support
	Storage Gateway(integrates On-premises IT environments with cloud storage)
		-> integrates with S3, Glacier, EBS
	Import/Export (Large Data Transfer) 
		-> support large data transfer, data transfered to from S3 or EBS
	Cloud Front (global content delivery network)
DataBases RDS, DynamoDb, ElastiCache, RedShift popular data base offerings
	RDS relational database
		-> resizable, backup, restore, 
		-> select from multiple engines like mysql, MS sql, oracle, postgress or 
		   Amazon Aurora (sits on top of mysql designed for cloud) 
	DynamoDb (predictable and scalable NoSQL Data Store)
		-> SSD drive solid state drive 
		-> runs on many node
	ElactiCache (In-Memory Cache)
		-> managed cached service
		-> support memcached and redis		
	RedShift (managed petabyte-Scale Data Warehouse)
		-> fully managed SQL based data warehouse
		-> ODBC/JDBC complient
		
Networking (VPC, ELB, Route 53, Direct Connect)
	VPC (virtual private cloud)
		-> lets you define cloud resourses,
		-> lets you define virtual cloud network, we can define subnets, 
		   Ip ranges, security groups, can be connected to datacenters over VPN or Direct Connect
		-> can be peered to another VPC
		-> managed by us
		-> creates user defined private cloud.
	Route 53(scalable DNS and Domain name registration and lookup)
		-> converts names into ip address
	ELB (Elastic load balancer)
		-> application load balances
		-> supports Http, https, Tcp traffic to EC2 instances
		-> dynamically grows and shrink based on traffic.
	Direct Connect(Dedicated network coonection to AWS)
		-> establishes a dedicated n/w connection from your premises to AWS
		-> reduces bandwith cost for high volume data transfer

Application Services(SES, SNS, SQS, Cloud search, Elastic transcoder, Simple Workflow service)
	SES (Simple email service) sending email services
		-> bulk and tranasctional email
		-> removes email server mgmt, meets ISP standard
		-> built in fedbak loop on succesfull & unsuccessful msg
	SNS (simple notification services) Push notification
		-> sends notication to mobile devices ioS, windows & andriod,
		-> simillar to sending to Pub server
	SQS (simple queue services) message queue service
		-> manages scalable message queue.
	Cloud search(manages cloud search)
		-> search functionality into application
	Elastic Transcoder(Easy-to-use media transcoder) 
		-> converts media files between convert
		-> integrates with S3 and Cloudfront
	Simple workflow service(workflow services for coordinating application components)
		-> coordinating processing steps across distributed system
		-> manages workflows including state, decision, executions, task and loging.

Adminstration and security (Directory services, IAM, Trusted advisor, CloudTrail, CloudWatch, AWS Config)
	Directory services (manages directory service in cloud)
		-> compatible with Active directory
	IAM (Identity Access and management)
		-> access control and key management to AWS infrastucture
		-> create users, groups, roles, 
		-> manages encrpytion keys
	Trusty advisor(AWS cloud optimization experts)
		-> checks AWS account for cost optimization, performance
		   fault tolerance and security, red flags
	CloudTrail (user activity and change tracking)
		-> Records AWS api call for your accounts, log files for Api call in S3
		   enables security analysis
	CloudWatch (Resources and application monitoring)
		-> visibility into resources utilization, operatonal performance
		-> supports custom metrics
		-> cpu, storage, threat count form dll etc
	AWS Config (resources configuring and management)		
		-> keeps history of onf changes..
		-> notification on configuring change

Deployment & management(Elastic bean, OpsWork, CloudFormation, CodeDeploy)
	Elastic Beanstalk (AWS application container)
		-> only upload their code and let services manages rest
		-> supports Docker, java. .Net, Node.js, Php and ruby
		-> automatically manages load balances.. etc.
	OpsWorks(Application management service)
		-> deploys, creates and automates application
		-> use chef recipies and cookbook
		-> prebuilt template for ruby, php, java and node.js
	CloudFormation (Templated AWS resources creation)
		-> model provisioning and update AWS in JSON formatted test files
		-> deploy stack from runtime parameters
	CodeDeploy (Automated deployments)
		-> automates deployments to EC2 instances 
		-> allows rolling updated and health checks
		-> use any language

Analytics(EMR)
	EMR(Elastic MapReduce) managed hadoop frameworks
		-> no h/w or virtuallization s/w
	Kinese (realtime processing of streaming big data) 
	Data pipeline( orchestration of data driven workflow)
	Machine learning (build smart application quickly and easily)
Enterprise Applications 	
	workspaces (Desktop in cloud)
	workdocs (Secure enterprize storage and sharing services)
	Workmail 



























	































	


			
