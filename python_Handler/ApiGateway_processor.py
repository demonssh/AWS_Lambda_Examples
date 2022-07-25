import json

print("Function loading")

def lambda_handler(event, context):

    trans_id = event['queryStringParameters']['trans_id']
    trans_type = event['queryStringParameters']['trans_type']
    trans_amount = event['queryStringParameters']['trans_amount']

    transactionResponse = {}
    transactionResponse['trans_id'] = trans_id
    transactionResponse['trans_type'] = trans_type
    transactionResponse['trans_amount'] = trans_amount
    transactionResponse['message'] = "Greeting!"

    responseObject = {}
    responseObject['statusCode'] = 200
    responseObject['body'] = json.dumps(transactionResponse)

    return responseObject

    #https://316y3onl6b.execute-api.us-east-2.amazonaws.com/test/transaction?trans_id=5&trans_type=purchase&trans_amount=2