opts = delimitedTextImportOptions("NumVariables", 18);

opts.DataLines = [1, Inf];
opts.Delimiter = " ";

opts.VariableNames = ["t", "V", "m", "h", "j", "d", "f", "X", "Xi", "K1", "Kp", "INa", "Isi", "Ik", "IK1", "IKP", "Ib", "stim"];
opts.VariableTypes = ["double", "double","double", "double","double", "double","double", "double","double","double","double", "double","double", "double","double", "double","double","double"];

opts.ExtraColumnsRule = "ignore";
opts.EmptyLineRule = "read";
opts.ConsecutiveDelimitersRule = "join";
opts.LeadingDelimitersRule = "ignore";

tbl = readtable("/Users/maxschool/go/src/goLR1/out.txt", opts);

t = tbl.t;
V = tbl.V;
m = tbl.m;
h = tbl.h;
j= tbl.j;
d = tbl.d;
f = tbl.f;
X = tbl.X;
Xi = tbl.Xi;
K1 = tbl.K1;
Kp = tbl.Kp;
INa = tbl.INa;
Isi = tbl.Isi;
Ik = tbl.Ik;
IK1 = tbl.IK1;
IKP = tbl.IKP;
Ib = tbl.Ib;
stim = tbl.stim;

%membrane voltage
figure(1)
sgtitle("LR1 Ventricular Action Potential", FontSize=16)
subplot(2,1,1)
plot(t,V, LineWidth=1.25)
xlim([min(t) max(t)])
grid(gca,'minor')
grid on;
xlabel("Time (ms)", FontSize=14)
ylabel("Membrane Voltage (mV)", FontSize=15)
legend('Vm', 'Location','best', 'FontSize',14)

subplot(2,1,2)
plot(t,stim, LineWidth=1.25)
xlim([min(t) max(t)])
grid(gca,'minor')
grid on;
xlabel("Time (ms)", FontSize=14)
ylabel("Current (\muA)", FontSize=15)
legend('Stimulus Injection', 'Location','best', 'FontSize',14)


%gates
figure(2)
sgtitle("Evolution of Gating Variables of LR1 Ventricular Action Potential", FontSize=16)
subplot(2,1,1)
plot(t,m,t,h,t,j,t,d,t,f,t,X,t,Xi,t,K1,t,Kp, LineWidth=1.25)
xlim([min(t) max(t)])
grid(gca,'minor')
grid on;
xlabel("Time (ms)", FontSize=14)
ylabel("Probability (unitless)", FontSize=15)
legend('m','h','j','d','f','X','Xi','K1','Kp', 'Location','best', 'FontSize',14)

subplot(2,1,2)
plot(t,stim, LineWidth=1.25)
xlim([min(t) max(t)])
grid(gca,'minor')
grid on;
xlabel("Time (ms)", FontSize=14)
ylabel("Current (\muA)", FontSize=15)
legend('Stimulus Injection', 'Location','best', 'FontSize',14)


%currents
figure(3)
sgtitle("Currents of LR1 Ventricular Action Potential", FontSize=16)
subplot(3,1,1)
plot(t,Isi,t,Ik,t,IK1,t,IKP,t,Ib, LineWidth=1.25)
xlim([min(t) max(t)])
grid(gca,'minor')
grid on;
xlabel("Time (ms)", FontSize=14)
ylabel("Current (\muA)", FontSize=15)
legend('Isi','Ik','IK1','IKP','Ib', 'Location','best', 'FontSize',14)

subplot(3,1,2)
plot(t,INa, LineWidth=1.25)
xlim([min(t) max(t)])
grid(gca,'minor')
grid on;
xlabel("Time (ms)", FontSize=14)
ylabel("Current (\muA)", FontSize=15)
legend('INa', 'Location','best', 'FontSize',14)

subplot(3,1,3)
plot(t,stim, LineWidth=1.25)
xlim([min(t) max(t)])
grid(gca,'minor')
grid on;
xlabel("Time (ms)", FontSize=14)
ylabel("Current (\muA)", FontSize=15)
legend('Stimulus Injection', 'Location','best', 'FontSize',14)
