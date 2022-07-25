import json
import boto3

dynamodb = boto3.resource('dynamodb')
table = dynamodb.Table('food')

def lambda_handler(event, context):
    # TODO implement
    print(table)
    taste = table.get_item(Key = {'name': 'KungPaoChicken'})
    print(taste)
    return {
        'statusCode': 200,
        'body': taste
    }