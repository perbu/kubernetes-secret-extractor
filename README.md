# Kubernetes Secret Extractor
## Description

This utility processes the output of kubectl get secret -o json <secret-name> commands, extracting the keys from the data section, decoding them from base64, and writing the results to files. It is designed to facilitate the handling of Kubernetes secrets in a local environment.

## Build & install

Clone the Repository:

```bash
git clone https://github.com/perbu/kubernetes-secret-extractor.git
cd kubernetes-secret-extractor
go build -o k8s-secret-extractor
```

## Usage

### Basic Command:
To use the utility, pipe the output of a kubectl get secret -o json command into it:

```bash
kubectl get secret -o json <secret-name> | ./k8s-secret-extractor
```

### Overwriting Existing Files:
If you want to overwrite existing files with the same names as the keys in the secret, use the --overwrite flag:

```bash
kubectl get secret -o json <secret-name> | ./k8s-secret-extractor --overwrite
```

### Security Note

This utility writes sensitive data to files on your local filesystem. Ensure that:

- Your local environment is secure.
- Files containing sensitive data are adequately protected.
- Secrets are handled in compliance with your organization's security policies.
