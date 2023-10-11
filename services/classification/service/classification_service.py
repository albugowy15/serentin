from proto import classification_pb2_grpc, classification_pb2
import os
import joblib
import pandas as pd

labels = ["Relaxed", "Calm", "Tense", "Stressed"];
BASE_DIR = os.path.dirname(os.path.abspath(__file__)) 

rf_model_path = os.path.join(BASE_DIR, 'model', 'random_forest_1.0.pkl')
rf_model = joblib.load(rf_model_path)

def predict(systolic_blood_pressure, diastolic_blood_pressure, heart_rate, body_temp):
    data = pd.DataFrame({
        'heart_rate': [heart_rate],
        'systolic_blood_pressure': [systolic_blood_pressure],
        'diastolic_blood_pressure': [diastolic_blood_pressure],
        'body_temp_c': [body_temp]
    })
    prediction  = rf_model.predict(data)
    stress_leve_index = int(prediction[0])
    return stress_leve_index

class ClassificationServicer(classification_pb2_grpc.ClassificationServicer):
    def PredictStressLevel(self, request, context):
        stress_level_index = predict(
            request.systolic_blood_pressure,
            request.diastolic_blood_pressure,
            request.heart_rate,
            request.body_temp_c
        )
        stress_level_label = labels[stress_level_index]
        return classification_pb2.StressLevelResponse(stress_level_index=stress_level_index, stress_level_label=stress_level_label)