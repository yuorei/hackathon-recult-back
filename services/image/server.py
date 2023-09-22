from concurrent import futures
import grpc
import image_pb2
import image_pb2_grpc
import os


class UploadImageServicer(image_pb2_grpc.UploadImageServicer):
    def Upload(self, request_iterator, context):
        image_data = b""
        meta = None
        count = 0
        for request in request_iterator:
            count += 1
            if request.WhichOneof('value') == 'meta':
                meta = request.meta

                print(str(count)+"回目(メタ情報)")
            elif request.WhichOneof('value') == 'data':
                image_data += request.data
                print(str(count)+"回目(画像データ)")

        # ファイルを保存するためのディレクトリを作成
        os.makedirs("image", exist_ok=True)

        # ファイルパスを生成
        file_path = os.path.join("image", meta.name)

        # 指定したフォルダに保存
        with open(file_path, "wb") as f:
            f.write(image_data)

        image_url = file_path

        response = image_pb2.UploadResponse(image_url=image_url)
        return response


def serve():
    port = "50052"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    image_pb2_grpc.add_UploadImageServicer_to_server(
        UploadImageServicer(), server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
