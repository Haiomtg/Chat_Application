# controllers/chat_controller.py
from flask import request, jsonify
from config.db import mongo

def send_message():
    data = request.get_json()
    user = data['user']
    text = data['text']

    mongo.db.messages.insert_one({'user': user, 'text': text})
    return jsonify(message='Message sent'), 201

def get_messages():
    messages = mongo.db.messages.find()
    return jsonify([{'user': msg['user'], 'text': msg['text']} for msg in messages])

