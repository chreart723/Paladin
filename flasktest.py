#!/bin/python

from flask import Flask, render_template
import requests 
import json

app = Flask(__name__)

def get_class(): 
    url = "https://www.dnd5eapi.co/api/classes/paladin"
    response = json.loads(requests.request("GET", url).text)
    dndclass = response["name"]
    return dndclass

@app.route("/")
def index():
    dndclass = get_class()
    # idpic_url = get_idpicture()
    return render_template("index.html", dndclass=dndclass)

# @app.route('/idpic')
# def get_image():
#     url = "https://enatwf3653.execute-api.us-east-1.amazonaws.com/default/92023-typaladin-getS3ImageObject"
#     response = json.loads(requests.request("GET", url).text)

#     # # response = json.loads(requests.request("GET", url).text)
#     # response = requests.get(url)
#     # base64_image_data_string = response["body"]
#     # base64_image = 
#     # response = Response(base64_image, content_type='image/png')
#     # return response
    

    