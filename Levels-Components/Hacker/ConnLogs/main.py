import random
import datetime



# List of possible source and destination IP addresses
source_ips = ['192.168.1.10', '192.168.1.20', '192.168.1.30']

# List of possible ports and protocols

# List of possible actions taken
actions = [
    {
        "summary": "The system has been restarted",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "A new user account has been created",
        "action": "User account credentials have been securely shared with the user",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "The firewall has been updated",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "A software update has been installed",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "A remote connection has been established",
        "action": "Connection has been verified and authorized",
        "protocol": "TCP",
        "port": "22"
    },
    {
        "summary": "An external device has been connected to the system",
        "action": "The device has been scanned and no threats were detected",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "The system has been shut down",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "An application has been installed",
        "action": "The application has been verified and authorized",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "The system has been backed up",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "A user has logged in",
        "action": "The login attempt has been verified and authorized",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "A user has logged out",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "A file has been deleted",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "The system clock has been updated",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "A security patch has been installed",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "A password has been reset",
        "action": "The new password has been securely shared with the user",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "A file has been downloaded",
        "action": "The file has been scanned and no threats were detected",
        "protocol": "TCP",
        "port": "443"
    },
    {
        "summary": "A file has been uploaded",
        "action": "The file has been scanned and no threats were detected",
        "protocol": "TCP",
        "port": "21"
    },
    {
        "summary": "The system has been updated",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "The system has been optimized",
        "action": "No action needed",
        "protocol": "",
        "port": ""
    },
     {
        "summary": "New user account created",
        "action": "Verified user account creation and permissions granted",
        "protocol": "",
        "port": ""
    },
    {
        "summary": "File transfer from external device",
        "action": "Scanned transferred file for malware",
        "protocol": "USB",
        "port": ""
    },
    {
        "summary": "SSH connection established",
        "action": "Verified user identity and purpose of connection",
        "protocol": "SSH",
        "port": "22"
    }
]


# Normal log report
abnormal_log_template = '''System Log Analysis Report

Date: {date}
Time: {time}

Summary:
The system has been compromised by an unauthorized user. The hacker gained access to the system using a backdoor that was created by exploiting a vulnerability in the system's firewall. The hacker used a remote IP address to gain access to the system.

Details:

    - Timestamp:         {timestamp}
    - Source IP Address: {source_ip}
    - Protocol:          {protocol}
    - Port:              {port}

    - Action Taken:
        - Administor alerted.

    - Notes:
        - User left a message : 
                                                             ..                ..                         s                                .                                          
                x=~                                        dF            . uW8"        ..                :8                               @88>                                        
 .d``          88x.   .e.   .e.     u.    u.              '88bu.         `t888        @L                .88       .u    .                 %8P      u.    u.                 .u    .   
 @8Ne.   .u   '8888X.x888:.x888   x@88k u@88c.      .u    '*88888bu       8888   .   9888i   .dL       :888ooo  .d88B :@8c        u        .     x@88k u@88c.      .u     .d88B :@8c  
 %8888:u@88N   `8888  888X '888k ^"8888""8888"   ud8888.    ^"*8888N      9888.z88N  `Y888k:*888.    -*8888888 ="8888f8888r    us888u.   .@88u  ^"8888""8888"   ud8888.  ="8888f8888r 
  `888I  888.   X888  888X  888X   8888  888R  :888'8888.  beWE "888L     9888  888E   888E  888I      8888      4888>'88"  .@88 "8888" ''888E`   8888  888R  :888'8888.   4888>'88"  
   888I  888I   X888  888X  888X   8888  888R  d888 '88%"  888E  888E     9888  888E   888E  888I      8888      4888> '    9888  9888    888E    8888  888R  d888 '88%"   4888> '    
   888I  888I   X888  888X  888X   8888  888R  8888.+"     888E  888E     9888  888E   888E  888I      8888      4888>      9888  9888    888E    8888  888R  8888.+"      4888>      
 uW888L  888'  .X888  888X. 888~   8888  888R  8888L       888E  888F     9888  888E   888E  888I     .8888Lu=  .d888L .+   9888  9888    888E    8888  888R  8888L       .d888L .+   
'*88888Nu88P   `%88%``"*888Y"     "*88*" 8888" '8888c. .+ .888N..888     .8888  888"  x888N><888'     ^%888*    ^"8888*"    9888  9888    888&   "*88*" 8888" '8888c. .+  ^"8888*"    
~ '88888F`       `~     `"          ""   'Y"    "88888%    `"888*""       `%888*%"     "88"  888        'Y"        "Y"      "888*""888"   R888"    ""   'Y"    "88888%       "Y"      
   888 ^                                          "YP'        ""             "`              88F                             ^Y"   ^Y'     ""                    "YP'                 
   *8E                                                                                      98"                                                                                       
   '8>                                                                                    ./"                                                                                         
    "                                                                                    ~`                                                                                           
'''

# Abnormal log report
log_template = '''System Log Analysis Report

Date: {date}
Time: {time}

Summary:
{summary}.

Details:

    - Timestamp:         {timestamp}
    - Source IP Address: {source_ip}
    - Protocol:          {protocol}
    - Port:              {port}
    - Action Taken:
        - {action_taken}.'''

start_date = datetime.datetime(2023, 1, 1)
end_date = datetime.datetime(2023, 5, 6)


# Generate 10 normal logs and 1 abnormal log
for i in range(500):
    timestamp = start_date + (end_date - start_date) * random.random()
    timestamp = timestamp.strftime("%Y-%m-%d %H:%M:%S")
    print(i, timestamp)
    if i == 5: # Generate the abnormal log
        source_ip = random.choice(source_ips)
        protocol = 'ssh'
        port = '22'
        log = abnormal_log_template.format(date=timestamp.split(" ")[0], time=timestamp.split(" ")[1], timestamp=timestamp, source_ip=source_ip, protocol=protocol, port=port)
        with open(f'temp/SPECIAL__{timestamp.replace(" ", "_")}.log', 'w') as f:
           f.write(log)
    else: # Gesnerate a normal log
        source_ip = random.choice(source_ips)
        action_taken = random.choice(actions)
        log = log_template.format(date=timestamp.split(" ")[0], time=timestamp.split(" ")[1], timestamp=timestamp, source_ip=source_ip, protocol=action_taken["protocol"], port=action_taken["port"], action_taken=action_taken["action"], summary=action_taken["summary"])
    # Save the log to a file
        with open(f'temp/{timestamp.replace(" ", "_")}.log', 'w') as f:
            f.write(log)
