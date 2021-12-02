import os
import sys

def main():
    if len(sys.argv) < 1:
        return

    table = '''
| Day | Prompt | Solution |
|---|---|---|
'''

    filelist = os.listdir(path=sys.argv[1])
    for file in filelist:
        table += f"| Day {file.strip('0')} | [Prompt]({sys.argv[1]}{file}/prompt) | [Solution]({sys.argv[1]}{file}) |\n"

    readme_template = open(sys.argv[2])
    readme = readme_template.read().replace("{{table}}", table)

    readme_template.close()
    f = open("README.md", "w")
    f.write(readme)
    f.close()

if __name__ == "__main__":
    main()
