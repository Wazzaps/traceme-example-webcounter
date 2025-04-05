FROM wazzaps/traceme:latest

ADD webcounter /webcounter

CMD ["/opt/traceme/bin/traceme", "/webcounter"]
