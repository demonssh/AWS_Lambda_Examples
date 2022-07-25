def lambda_handler(event, context):
    message = 'Hello {} {}'.format(event['first_name'], event['last_name'])
    return {
        'message' : message
    }



#test
# {
#   "first_name": "John",
#   "last_name": "Bob"
# }