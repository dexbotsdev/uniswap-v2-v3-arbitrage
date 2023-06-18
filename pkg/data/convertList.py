import json

# Open the whitelist.go file for reading
with open("whitelist.go", "r") as f:
    # Read the contents of the file into a string
    file_contents = f.read()

    # Find the start and end of the Whitelist array
    start_index = file_contents.find("Whitelist = [") + len("Whitelist = [")
    end_index = file_contents.find("]", start_index)

    # Extract the Whitelist array from the string and split it into a list
    whitelist_str = file_contents[start_index:end_index]
    whitelist_list = whitelist_str.strip().strip(",").split("\n")

    # Strip any whitespace and quotes from each element in the list
    whitelist = [element.strip().strip('"') for element in whitelist_list]

# Write the Whitelist array to a JSON file
with open("whitelist.json", "w") as f:
    json.dump(whitelist, f)