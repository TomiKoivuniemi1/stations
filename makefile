.PHONY: all

all: test1 test2 test3 test4 test5 test6 test7 test8 test9 test10 \
     test11 test12 test13 test14 test15 test16 test17 test18 test19 \
     test20 test21 test22 test23 test24 test25 test26 test27 test28 \
     test29

test1:
	@echo "Running test 1: go run . testing/duplicateconnections.map beginning terminus 20" >> testresults.txt
	-@go run . testing/duplicateconnections.map beginning terminus 20 >> testresults.txt 2>&1

test2:
	@echo "Running test 2: go run . testing/network6.map waterloo st_pancras 3" >> testresults.txt
	-@go run . testing/network6.map waterloo st_pancras 3 >> testresults.txt 2>&1

test3:
	@echo "Running test 3: go run . testing/network6.map waterloo st_pancras 3 too many" >> testresults.txt
	-@go run . testing/network6.map waterloo st_pancras 3 too many >> testresults.txt 2>&1

test4:
	@echo "Running test 4: go run . testing/noconnections.map beginning terminus 20" >> testresults.txt
	-@go run . testing/noconnections.map beginning terminus 20 >> testresults.txt 2>&1

test5:
	@echo "Running test 5: go run . testing/network.map beethoven part 9" >> testresults.txt
	-@go run . testing/network.map beethoven part 9 >> testresults.txt 2>&1

test6:
	@echo "Running test 6: go run . testing/doublestations.map beginning terminus 20" >> testresults.txt
	-@go run . testing/doublestations.map beginning terminus 20 >> testresults.txt 2>&1

test7:
	@echo "Running test 7: go run . testing/network6.map too few" >> testresults.txt
	-@go run . testing/network6.map too few >> testresults.txt 2>&1

test8:
	@echo "Running test 8: go run . testing/network.map startstationmissing part 9" >> testresults.txt
	-@go run . testing/network.map startstationmissing part 9 >> testresults.txt 2>&1

test9:
	@echo "Running test 9: go run . testing/network.map beethoven endstationmissing 9" >> testresults.txt
	-@go run . testing/network.map beethoven endstationmissing 9 >> testresults.txt 2>&1

test10:
	@echo "Running test 10: go run . testing/network.map beethoven part 9" >> testresults.txt
	-@go run . testing/network.map beethoven part 9 >> testresults.txt 2>&1

test11:
	@echo "Running test 11: go run . testing/stationnotexist.map beginning terminus 20" >> testresults.txt
	-@go run . testing/stationnotexist.map beginning terminus 20 >> testresults.txt 2>&1

test12:
	@echo "Running test 12: go run . testing/nostations.map beginning terminus 20" >> testresults.txt
	-@go run . testing/nostations.map beginning terminus 20 >> testresults.txt 2>&1

test13:
	@echo "Running test 13: go run . testing/coordinates.map beginning terminus 20" >> testresults.txt
	-@go run . testing/coordinates.map beginning terminus 20 >> testresults.txt 2>&1

test14:
	@echo "Running test 14: go run . testing/network2.map small large 9" >> testresults.txt
	-@go run . testing/network2.map small large 9 >> testresults.txt 2>&1

test15:
	@echo "Running test 15: go run . testing/network3.map two four 4" >> testresults.txt
	-@go run . testing/network3.map two four 4 >> testresults.txt 2>&1

test16:
	@echo "Running test 16: go run . testing/network4.map jungle desert 10" >> testresults.txt
	-@go run . testing/network4.map jungle desert 10 >> testresults.txt 2>&1

test17:
	@echo "Running test 17: go run . testing/network7.map beginning terminus -20" >> testresults.txt
	-@go run . testing/network7.map beginning terminus -20 >> testresults.txt 2>&1

test18:
	@echo "Running test 18: go run . testing/network6.map waterloo st_pancras 100" >> testresults.txt
	-@go run . testing/network6.map waterloo st_pancras 100 >> testresults.txt 2>&1

test19:
	@echo "Running test 19: go run . testing/invalidnames.map beginning terminus 20" >> testresults.txt
	-@go run . testing/invalidnames.map beginning terminus 20 >> testresults.txt 2>&1

test20:
	@echo "Running test 20: go run . testing/network5.map bond_square space_port 4" >> testresults.txt
	-@go run . testing/network5.map bond_square space_port 4 >> testresults.txt 2>&1

test21:
	@echo "Running test 21: go run . testing/samecoordinates.map beginning terminus 20" >> testresults.txt
	-@go run . testing/samecoordinates.map beginning terminus 20 >> testresults.txt 2>&1

test22:
	@echo "Running test 22: go run . testing/mandatoryfix.map waterloo st_pancras 4" >> testresults.txt
	-@go run . testing/mandatoryfix.map waterloo st_pancras 4 >> testresults.txt 2>&1

test23:
	@echo "Running test 23: go run . testing/network6.map waterloo st_pancras 2" >> testresults.txt
	-@go run . testing/network6.map waterloo st_pancras 2 >> testresults.txt 2>&1

test24:
	@echo "Running test 24: go run . testing/network7.map beginning terminus 20" >> testresults.txt
	-@go run . testing/network7.map beginning terminus 20 >> testresults.txt 2>&1

test25:
	@echo "Running test 25: go run . testing/network6.map waterloo st_pancras 1" >> testresults.txt
	-@go run . testing/network6.map waterloo st_pancras 1 >> testresults.txt 2>&1

test26:
	@echo "Running test 26: go run . testing/network6.map waterloo st_pancras 4" >> testresults.txt
	-@go run . testing/network6.map waterloo st_pancras 4 >> testresults.txt 2>&1

test27:
	@echo "Running test 27: go run . testing/network6.map waterloo waterloo 4" >> testresults.txt
	-@go run . testing/network6.map waterloo waterloo 4 >> testresults.txt 2>&1

test28:
	@echo "Running test 28: go run . testing/nopath.map beginning terminus 20" >> testresults.txt
	-@go run . testing/nopath.map beginning terminus 20 >> testresults.txt 2>&1

test29:
	@echo "Running test 29: go run . testing/tenthousand.map station_00000 station_00001 1" >> testresults.txt
	-@go run . testing/tenthousand.map station_00000 station_00001 1 >> testresults.txt 2>&1
