import kagglehub

# Download latest version
path = kagglehub.dataset_download("lydia70/education-in-africa")

print("Path to dataset files:", path)