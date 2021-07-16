
echo "kill process ..."

ssh -i "big.pem" ubuntu@ec2-54-174-95-219.compute-1.amazonaws.com  "pidof api-wecode-supplychain | awk '{print $1}' | xargs kill "
