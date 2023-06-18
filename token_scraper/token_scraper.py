from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.common.by import By
from web3 import Web3
import time

# Setup Selenium Chrome WebDriver
driver = webdriver.Chrome(executable_path=ChromeDriverManager().install())

# List to store token addresses
token_addresses = []

# Base URL
base_url = 'https://etherscan.io'

# Iterate over the first 26 pages
for i in range(1, 27):
    driver.get(f'https://etherscan.io/tokens?p={i}')
    
    # Allow the page to load
    time.sleep(3)

    # Find all 'a' tags with the specific class
    a_tags = driver.find_elements(By.CSS_SELECTOR, 'a.d-flex.align-items-center.gap-1.link-dark')

    for tag in a_tags:
        # Extract href attribute and append to token_addresses
        token_addresses.append(tag.get_attribute('href'))

# Close the browser
driver.quit()

# List to store checksummed addresses
checksum_addresses = []

# Get all token addresses
for url in token_addresses:
    # Get address from url
    address = url.split('/')[-1]
    
    # Checksum the address
    checksum_address = Web3.toChecksumAddress(address)
    
    # Append checksummed address to list
    checksum_addresses.append(checksum_address)

# Write checksummed addresses to file as comma-separated
with open('out.txt', 'w') as f:
    for address in checksum_addresses:
        f.write('"' + address + '",\n')
