from app import app
import requests
from flask import request, jsonify
import asyncio
import os
# @app.get('/plot', methods=['GET'])
# def plotData():
    
#     headers = {
#         'api_key': os.getenv()['API_KEY'],
#         'data': ['value'],
#         'length':500
#     }
#     req = requests.get('https://api.eia.gov/v2/co2-emissions/co2-emissions-aggregates/data/?frequency=annual', headers=headers)
#     print(req)