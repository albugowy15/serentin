from concurrent import futures
import grpc
import logging
from dotenv import load_dotenv
from proto import classification_pb2_grpc
from service.classification_service import ClassificationServicer

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    classification_pb2_grpc.add_ClassificationServicer_to_server(
        ClassificationServicer(), server
    )
    server.add_insecure_port("[::]:50052")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    load_dotenv()
    serve()  # Pass the db connection to the serve function
