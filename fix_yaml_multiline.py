import sys
from pathlib import Path
from ruamel.yaml import YAML
from ruamel.yaml.scalarstring import PlainScalarString, SingleQuotedScalarString, FoldedScalarString, LiteralScalarString

def convert_to_literal_if_multiline(data_node):
    """
    Recursively traverses the YAML data structure.
    If a string node (plain, single-quoted, or folded) contains newlines,
    it's converted to a LiteralScalarString.
    """
    if isinstance(data_node, dict):
        # Iterate over a copy of items for safe modification if necessary (though we are returning new values)
        for key, value in data_node.items():
            data_node[key] = convert_to_literal_if_multiline(value)
    elif isinstance(data_node, list):
        for i, item in enumerate(data_node):
            data_node[i] = convert_to_literal_if_multiline(item)
    elif isinstance(data_node, (PlainScalarString, SingleQuotedScalarString, FoldedScalarString)):
        # Check if the string's actual value contains a newline character
        if '\n' in str(data_node): # Use str(data_node) to get the value
            # If it's not already a LiteralScalarString, convert it.
            # This preserves the original content, but changes its YAML representation.
            return LiteralScalarString(str(data_node))
    elif isinstance(data_node, str): # Handle raw Python strings if they appear
        if '\n' in data_node:
            return LiteralScalarString(data_node)
    
    # If it's already a LiteralScalarString or any other type, return as is.
    return data_node

def main():
    if len(sys.argv) < 2:
        print("Usage: python fix_yaml_multiline.py <input_yaml_file> [output_yaml_file]")
        print("If output_yaml_file is not provided, it will be <input_yaml_file_stem>_fixed.yaml")
        sys.exit(1)

    input_file_path = Path(sys.argv[1])
    if len(sys.argv) > 2:
        output_file_path = Path(sys.argv[2])
    else:
        output_file_path = input_file_path.parent / f"{input_file_path.stem}_fixed{input_file_path.suffix}"

    if not input_file_path.is_file():
        print(f"Error: Input file '{input_file_path}' not found.")
        sys.exit(1)

    yaml = YAML()
    yaml.preserve_quotes = True  # Tries to preserve existing quotes where not overridden
    yaml.indent(mapping=2, sequence=4, offset=2) # Common indentation, adjust if needed

    print(f"Loading YAML from: {input_file_path}")
    try:
        with input_file_path.open('r', encoding='utf-8') as f:
            data = yaml.load(f)
        print("YAML loaded successfully.")
    except Exception as e:
        print(f"Error: Could not parse YAML file '{input_file_path}'.")
        print(f"Details: {e}")
        print("The script cannot fix files that are too malformed to be parsed by ruamel.yaml.")
        print("Please try to fix the most critical syntax errors manually first,")
        print(f"especially the one indicated around line 221 in your original error message.")
        sys.exit(1)

    print("Processing YAML to convert multi-line strings to literal block style...")
    fixed_data = convert_to_literal_if_multiline(data)

    print(f"Saving modified YAML to: {output_file_path}")
    try:
        with output_file_path.open('w', encoding='utf-8') as f:
            yaml.dump(fixed_data, f)
        print("YAML processing complete. Output saved.")
        print(f"Please review the changes in '{output_file_path}'.")
    except Exception as e:
        print(f"Error: Could not write YAML to file '{output_file_path}'.")
        print(f"Details: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()
