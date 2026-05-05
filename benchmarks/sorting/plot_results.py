import subprocess
import matplotlib.pyplot as plt
import os

BASE_DIR = os.path.dirname(os.path.abspath(__file__))

# пути к файлам
python_file = os.path.join(BASE_DIR, "python_benchmark.py")
go_file = os.path.join(BASE_DIR, "go_benchmark.go")

py_result = subprocess.check_output(["python", python_file]).decode()
go_result = subprocess.check_output(["go", "run", go_file]).decode()

def parse_time(output):
    return float(output.strip().split(":")[1])

python_time = parse_time(py_result)
go_time = parse_time(go_result)

languages = ["Python", "Go"]
times = [python_time, go_time]

plt.bar(languages, times)
plt.title("Automatic Benchmark: Python vs Go")
plt.ylabel("Time (seconds)")

plt.savefig(os.path.join(BASE_DIR, "benchmark_result.png"))