# Libraries
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

def plot(list, list1, users, title, label):
    barWidth = 0.25

    r1 = np.arange(len(list))
    r2 = [x + barWidth for x in r1]

    plt.xlabel('Número de Clientes')

    plt.ylabel('Tempo ('+label+')')
    plt.title(title+' (N = 10000)')

    plt.bar(r1, list1, color='#13293D', width=barWidth, edgecolor='white', label='UDP') #13293D #006494 #FF7F00
    plt.bar(r2, list, color='#70C1B3', width=barWidth, edgecolor='white', label='TCP') #5D707F #70C1B3

    # Create names on the x-axis
    plt.xticks([r + barWidth for r in range(len(list))], users)

    # Show graphic
    plt.legend()
    plt.show()

def CalcMeanTime(path, users, type):
    listMean = []
    listTot = []
    dataframe = pd.DataFrame()
    print(type)
    for i in users:
        #valores do tempo em ns
        data = np.genfromtxt(fname=path+ str(i) +'.txt')
        print("users = ", i,"| Std(micros):" ,round(np.std(data)/1000), "| Mean(micros):" ,round(np.mean(data)/1000), "|Tempo Total (ms):", round(np.sum(data)/10e6))
        listMean.append(np.mean(data)/1000)
        listTot.append(np.sum(data)/10e6)
        dataframe = pd.concat([dataframe, makeDf(data, i)])

    return listMean, listTot, dataframe

def makeDf (data, user):
    NumUser = []
    for i in range(int(len(data))):
        NumUser.append(int(user)) 
    df = pd.DataFrame(list(zip(NumUser,data)), columns = ['User','Time'])
    NumUser.clear()
    return df

def BoxPlot(dataframe, title):
    boxplot = dataframe.boxplot(column=['Time'],by="User",grid=False, rot=45, fontsize=15)
    plt.title('')
    plt.suptitle('')
    boxplot.set_title(title);
    boxplot.set_xlabel("Número de Users");
    boxplot.set_ylabel("Tempo (ns)");
    boxplot.plot()
    boxplot = plt.show()

users = [1, 2, 5, 10]

#dataframe
TCP=pd.DataFrame()
UDP=pd.DataFrame()

#Mean Time Array, Total Time
TCPmean, TCPtot,TCP = CalcMeanTime('/home/katarine/Documents/Mestrado/Segundo_Semestre/Plataformas/codes/Projects/ConsultaLocal/socketTCP/client/time01_', users, 'TCP')
UDPmean, UDPtot, UDP = CalcMeanTime('/home/katarine/Documents/Mestrado/Segundo_Semestre/Plataformas/codes/Projects/ConsultaLocal/socketUDP/client/time01_', users, 'UDP')

plot(TCPmean, UDPmean, users,'Tempo Médio','micros')
plot(TCPtot, UDPtot, users, 'Tempo Total', 'ms')

print(TCP.describe())
print(UDP.describe())

BoxPlot(TCP, 'TCP')
BoxPlot(UDP, 'UDP')
