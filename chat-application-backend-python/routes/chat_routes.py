# routes/chat_routes.py
from flask import Blueprint
from controllers.chat_controller import send_message, get_messages
from middleware.auth_middleware import auth_middleware

chat_bp = Blueprint('chat', __name__)

chat_bp.route('/messages', methods=['POST'])(auth_middleware)(send_message)
chat_bp.route('/messages', methods=['GET'])(auth_middleware)(get_messages)

