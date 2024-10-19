# app.py
from flask import Flask
from config.db import mongo
from routes.auth_routes import auth_bp
from routes.chat_routes import chat_bp
from flask_jwt_extended import JWTManager

app = Flask(__name__)
app.config["MONGO_URI"] = "mongodb://localhost:27017/chat-app"
app.config["JWT_SECRET_KEY"] = "your_jwt_secret"

mongo.init_app(app)
jwt = JWTManager(app)

app.register_blueprint(auth_bp, url_prefix='/api/auth')
app.register_blueprint(chat_bp, url_prefix='/api/chat')

if __name__ == '__main__':
    app.run(port=5000)

