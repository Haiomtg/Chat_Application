# controllers/auth_controller.py
from flask import request, jsonify
from flask_jwt_extended import create_access_token
from werkzeug.security import generate_password_hash, check_password_hash
from config.db import mongo

def register():
    data = request.get_json()
    username = data['username']
    password = generate_password_hash(data['password'])

    mongo.db.users.insert_one({'username': username, 'password': password})
    return jsonify(message='User registered'), 201

def login():
    data = request.get_json()
    username = data['username']
    password = data['password']

    user = mongo.db.users.find_one({'username': username})
    if user and check_password_hash(user['password'], password):
        access_token = create_access_token(identity=username)
        return jsonify(access_token=access_token)

    return jsonify(message='Invalid credentials'), 401

