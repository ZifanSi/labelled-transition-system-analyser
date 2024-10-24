import java.awt.*;
import java.awt.event.*;
import javax.swing.*;

public class FMRadio extends JFrame implements ActionListener {
    private JLabel label;
    private JLabel frequencyDisplay;
    private JButton onOffBtn, scanBtn, resetBtn, lockBtn;
    private JTextArea guide;
    private boolean isOn;
    private boolean isLocked;
    private int frequency;
    private Timer scanTimer;

    private final int TOP_FREQUENCY = 108;
    private final int BOTTOM_FREQUENCY = 88;
    private final int STEP = 1;
    
    public FMRadio() {
        setTitle("FM Radio");
        setLayout(new BorderLayout());

        isOn = false;
        isLocked = false;
        frequency = TOP_FREQUENCY;

        label = new JLabel("Radio is OFF");
        frequencyDisplay = new JLabel("Frequency: -- MHz");

        onOffBtn = new JButton("On/Off");
        scanBtn = new JButton("Scan");
        resetBtn = new JButton("Reset");
        lockBtn = new JButton("Lock");

        scanBtn.setEnabled(false);
        resetBtn.setEnabled(false);
        lockBtn.setEnabled(false);

        onOffBtn.addActionListener(this);
        scanBtn.addActionListener(this);
        resetBtn.addActionListener(this);
        lockBtn.addActionListener(this);

        JPanel controlPanel = new JPanel();
        controlPanel.setLayout(new GridLayout(3, 2, 5, 5));

        controlPanel.add(onOffBtn);
        controlPanel.add(label);
        controlPanel.add(scanBtn);
        controlPanel.add(resetBtn);
        controlPanel.add(lockBtn);
        controlPanel.add(frequencyDisplay);

        guide = new JTextArea(5, 20);
        guide.setText("HOW TO USE FM Radio ???\n"
            + "1. On/Off: Turns the radio on/off.\n"
            + "2. Scan: Scans from 108 MHz to 88 MHz.\n"
            + "3. Reset: Resets the frequency to 108 MHz.\n"
            + "4. Lock: Locks the radio at the current frequency.\n"
            + "5. Click again to unlock, and click scan can resume after unlocking.\n");

        guide.setEditable(false);
        guide.setLineWrap(true);
        guide.setWrapStyleWord(true);

        JScrollPane scrollPane = new JScrollPane(guide);

        add(controlPanel, BorderLayout.CENTER);
        add(scrollPane, BorderLayout.SOUTH);

        scanTimer = new Timer(500, new ActionListener() {
            public void actionPerformed(ActionEvent e) {
                scanFrequency();
            }
        });

        setSize(500, 500);
        setLocationRelativeTo(null);
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setVisible(true);
    }

    public void actionPerformed(ActionEvent ae) {
        if (ae.getSource() == onOffBtn) {
            toggleRadio();
        } else if (ae.getSource() == scanBtn) {
            startScanning();
        } else if (ae.getSource() == resetBtn) {
            resetFrequency();
        } else if (ae.getSource() == lockBtn) {
            toggleLock();
        }
    }

    private void toggleRadio() {
        if (!isOn) {
            isOn = true;
            isLocked = false;
            frequency = TOP_FREQUENCY;
            label.setText("Radio is ON");
            frequencyDisplay.setText("Frequency: " + frequency + " MHz");
            scanBtn.setEnabled(true);
            resetBtn.setEnabled(true);
            lockBtn.setEnabled(true);
        } else {
            isOn = false;
            label.setText("Radio is OFF");
            frequencyDisplay.setText("Frequency: -- MHz");
            scanBtn.setEnabled(false);
            resetBtn.setEnabled(false);
            lockBtn.setEnabled(false);
            scanTimer.stop();
        }
    }

    private void startScanning() {
        if (isOn && !scanTimer.isRunning() && !isLocked) {
            scanTimer.start();
        }
    }

    private void scanFrequency() {
        if (frequency > BOTTOM_FREQUENCY) {
            frequency -= STEP;
            frequencyDisplay.setText("Frequency: " + frequency + " MHz");
        } else {
            scanTimer.stop();
        }
    }

    private void resetFrequency() {
        if (isOn) {
            isLocked = false;
            frequency = TOP_FREQUENCY;
            frequencyDisplay.setText("Frequency: " + frequency + " MHz");
            scanTimer.stop();
            label.setText("Radio is ON");
        }
    }

    private void toggleLock() {
        if (isOn) {
            if (isLocked) {
                isLocked = false;
                label.setText("Radio is UNLOCKED");
            } else {
                isLocked = true;
                scanTimer.stop();
                label.setText("Radio is LOCKED at " + frequency + " MHz");
            }
        }
    }

    public static void main(String[] args) {
        new FMRadio();
    }
}
