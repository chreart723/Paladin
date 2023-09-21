import base64
import json
import boto3

def lambda_handler(event, context):
    bucket_name = "92023-typaladin-dndapp"
    s3_obj_key = "/images/ty.jpg"
    s3_client = boto3.client('s3', 'us-east-1')

    read_from_s3 = s3_client.get_object(
        Key=s3_obj_key
        Bucket=bucket_name
    )
    image = response['Body'].read()
    return { 
        'headers': {"Content-Type": "image/png"},
        'statusCode': 200,
        'body': base64.b64encode(image),
        'isBase64Encoded': True
    }