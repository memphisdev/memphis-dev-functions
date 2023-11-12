# Example Function: Adding A Severity Field

This function's goal is to check a field, in this case, <time_since_last_produce\> and add a json field <severity\> that could be used to check how well a given service is behaving. 

## Example Use Case Definition

A user has a service which they cannot change, but has data in their messages that is a quantative measure which could possibly be used to check if a service is being overloaded or not. 

With Memphis Functions, the user could simply attach a function to the station (or to a separate station which a copy of every kth message might be sent to in order to save on Lambda compute time). This function could check that qualitive measure and change the message that is being produced to warn that the service is being overloaded. Or, even better, the Function could directly alert a monitoring tool of the health of the producing service. 

This could be used to monitor the health of the system and to make sure that latency or other measures is optimal.