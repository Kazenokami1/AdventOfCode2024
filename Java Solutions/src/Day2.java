/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

/**
 *
 * @author dhc10
 */
public class Day2 {
	
    public static void main(String args[]) {
        Integer safeReports = 0;
        ArrayList<ArrayList<Integer>> reports = new ArrayList<ArrayList<Integer>>();
         try {
            File lists = new File("../inputFiles/Day2.txt");
            Scanner myReader = new Scanner(lists);
            while(myReader.hasNextLine()) {
                String[] levelStrings = myReader.nextLine().split(" ");
                ArrayList<Integer> levels = new ArrayList<Integer>();
                for (int i=0;i<levelStrings.length;i++) {
                    levels.add(Integer.parseInt(levelStrings[i]));
                }
                reports.add(levels);
            }
            myReader.close();
        } catch (FileNotFoundException e) {
            System.out.println("No File Found");
            e.printStackTrace();
        } 
        for (ArrayList<Integer> report : reports) {
            if (isListSorted(report)){
                if (checkReportSafety(report)) {
                    safeReports++;
                }
            }
        }
        System.out.printf("Part 1 Answer: %d\n", safeReports);
        safeReports = 0;
        for (ArrayList<Integer> report : reports) {
            for (int i=0;i<report.size();i++) {
                ArrayList<Integer> newReport = new ArrayList<Integer>(report);
                newReport.remove(i);
                if (isListSorted(newReport)) {
                    if (checkReportSafety(newReport)) {
                        safeReports++;
                        break;
                    } 
                }
            }
        }
        System.out.printf("Part 2 Answer: %d\n", safeReports);
    }

    private static boolean checkReportSafety(ArrayList<Integer> report) {
        for (int i=0;i<report.size()-1;i++) {
            if (Math.abs(report.get(i)-report.get(i+1)) > 3 || report.get(i)==report.get(i+1)) {
                return false;
            }
        }
        return true;
    }

    private static boolean isListSorted(ArrayList<Integer> report) {
        Boolean listSorted = true;
        for (int i = 0;i<report.size()-1;i++) {
            if (report.get(i) > report.get(i+1)) {
                listSorted = false;
                break;
            }
        }
        if (!listSorted) {
            listSorted = true;
            for (int i=0;i<report.size()-1;i++) {
                if (report.get(i)<report.get(i+1)) {
                    listSorted = false;
                    break;
                }
            }
        }
        return listSorted;
    }
}